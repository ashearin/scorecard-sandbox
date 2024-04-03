// Copyright 2024 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package githubrepo

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"sync"

	"github.com/google/go-github/v59/github"

	"github.com/ossf/scorecard/v4/clients"
)

type sbomHandler struct {
	ghclient *github.Client
	once     *sync.Once
	ctx      context.Context
	errSetup error
	repourl  *repoURL
	sboms    []clients.Sbom
}

func (handler *sbomHandler) init(ctx context.Context, repourl *repoURL) {
	handler.ctx = ctx
	handler.repourl = repourl
	handler.errSetup = nil
	handler.once = new(sync.Once)
	handler.sboms = nil
}

func (handler *sbomHandler) setup() error {
	handler.once.Do(func() {
		// Check for sboms in release artifacts
		err := handler.checkReleaseArtifacts()
		if err != nil {
			handler.errSetup = fmt.Errorf("failed searching for Sbom in Release artifacts: %w", err)
		}

		// Check for sboms in pipeline artifacts
		err = handler.checkCICDArtifacts()
		if err != nil {
			handler.errSetup = fmt.Errorf("failed searching for Sbom in CICD artifacts: %w", err)
		}

		// Hit Github api and grab autogenerated sbom
		err = handler.fetchGithubAPISbom()
		if err != nil {
			handler.errSetup = fmt.Errorf("failed retrieving github autogenerated sbom: %w", err)
		}
	})

	return handler.errSetup
}

func (handler *sbomHandler) listSboms() ([]clients.Sbom, error) {
	if err := handler.setup(); err != nil {
		return nil, fmt.Errorf("error during sbomHandler.setup: %w", err)
	}

	return handler.sboms, nil
}

func (handler *sbomHandler) checkReleaseArtifacts() error {
	client := handler.ghclient

	// defined at: (using apiVersion=2022-11-28)
	// docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#get-the-latest-release
	reqURL := path.Join("repos", handler.repourl.owner, handler.repourl.repo, "releases", "latest")
	req, err := client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("request for repo latest release failed with %w", err)
	}
	bodyJSON := github.RepositoryRelease{}
	// The client.repoClient.Do API writes the response body to var bodyJSON,
	// so we can ignore the first returned variable (the entire http response object)
	// since we only need the response body here.
	resp, derr := client.Do(handler.ctx, req, &bodyJSON)
	if derr != nil {
		return fmt.Errorf("response for repo latest release failed with %w", derr)
	}
	if resp.StatusCode != http.StatusOK {
		// Dont fail, just return
		// TODO: print info for users that a non-200 response was returned
		return nil
	}

	if len(bodyJSON.Assets) > 0 {
		for i := range bodyJSON.Assets {
			asset := bodyJSON.Assets[i]

			if !clients.ReSbomFile.MatchString(asset.GetName()) {
				continue
			}

			handler.sboms = append(handler.sboms, clients.Sbom{
				Name:    asset.GetName(),
				Origin:  "repositoryRelease",
				URL:     asset.GetBrowserDownloadURL(),
				Created: asset.CreatedAt.Time,
				Path:    asset.GetURL(),
			},
			)
		}
	}

	return nil
}

func (handler *sbomHandler) checkCICDArtifacts() error {
	// Originally wanted to use workflowruns from latest release, but
	// that would've resulted in as many api calls as workflows runs (11 for scorcard itself)
	// Seems like deficiency in github api (or my understanding of it)
	client := handler.ghclient
	// defined at: (using apiVersion=2022-11-28)
	// docs.github.com/en/rest/actions/artifacts?#list-artifacts-for-a-repository
	reqURL := path.Join("repos", handler.repourl.owner, handler.repourl.repo, "actions", "artifacts")
	req, err := client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("request for repo artifacts failed with %w", err)
	}
	bodyJSON := github.ArtifactList{}
	// The client.repoClient.Do API writes the response body to var bodyJSON,
	// so we can ignore the first returned variable (the entire http response object)
	// since we only need the response body here.
	resp, derr := client.Do(handler.ctx, req, &bodyJSON)
	if derr != nil {
		return fmt.Errorf("response for repo sbom failed with %w", derr)
	}

	if resp.StatusCode != http.StatusOK {
		// Dont fail, just return
		// TODO: print info for users that a non-200 response was returned
		return nil
	}

	if bodyJSON.GetTotalCount() == 0 {
		return nil
	}

	returnedArtifacts := bodyJSON.Artifacts

	for i := range returnedArtifacts {
		artifact := returnedArtifacts[i]

		if *artifact.Expired || !clients.ReSbomFile.MatchString(artifact.GetName()) {
			continue
		}

		handler.sboms = append(handler.sboms, clients.Sbom{
			Name:    artifact.GetName(),
			Origin:  "repositoryCICD",
			URL:     artifact.GetArchiveDownloadURL(),
			Created: artifact.CreatedAt.Time,
			Path:    artifact.GetURL(),
		},
		)
	}

	return nil
}

func (handler *sbomHandler) fetchGithubAPISbom() error {
	client := handler.ghclient
	// defined at: (using apiVersion=2022-11-28)
	// docs.github.com/en/rest/dependency-graph/sboms#export-a-software-bill-of-materials-sbom-for-a-repository
	reqURL := path.Join("repos", handler.repourl.owner, handler.repourl.repo, "dependency-graph", "sbom")
	req, err := client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("request for repo sbom failed with %w", err)
	}
	bodyJSON := github.SBOM{}
	// The client.repoClient.Do API writes the response body to var bodyJSON,
	// so we can ignore the first returned variable (the entire http response object)
	// since we only need the response body here.
	resp, derr := client.Do(handler.ctx, req, &bodyJSON)
	if derr != nil {
		return fmt.Errorf("response for repo sbom failed with %w", derr)
	}
	if resp.StatusCode != http.StatusOK {
		// Dont fail, just return
		return nil
	}

	ReturnedSbom := bodyJSON.GetSBOM()

	if ReturnedSbom != nil {
		handler.sboms = append(handler.sboms, clients.Sbom{
			ID:            *ReturnedSbom.SPDXID,
			Name:          *ReturnedSbom.Name,
			Origin:        "repositoryRelease",
			URL:           *ReturnedSbom.DocumentNamespace,
			Created:       ReturnedSbom.CreationInfo.Created.Time,
			Path:          ReturnedSbom.DocumentDescribes[0],
			Tool:          ReturnedSbom.CreationInfo.Creators[0],
			Schema:        "SPDX",
			SchemaVersion: *ReturnedSbom.SPDXVersion,
		},
		)
	}
	return nil
}
