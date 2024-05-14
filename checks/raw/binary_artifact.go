// Copyright 2021 OpenSSF Scorecard Authors
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

package raw

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/rhysd/actionlint"

	"github.com/ossf/scorecard/v5/checker"
	"github.com/ossf/scorecard/v5/checks/fileparser"
	"github.com/ossf/scorecard/v5/clients"
	sce "github.com/ossf/scorecard/v5/errors"
	"github.com/ossf/scorecard/v5/finding"
)

// how many bytes are considered when determining if a file is text or binary.
const binaryTestLen = 1024

// BinaryArtifacts retrieves the raw data for the Binary-Artifacts check.
func BinaryArtifacts(req *checker.CheckRequest) (checker.BinaryArtifactData, error) {
	c := req.RepoClient
	files := []checker.File{}
	err := fileparser.OnMatchingFileReaderDo(c, fileparser.PathMatcher{
		Pattern:       "*",
		CaseSensitive: false,
	}, checkBinaryFileReader, &files)
	if err != nil {
		return checker.BinaryArtifactData{}, fmt.Errorf("%w", err)
	}
	// Ignore validated gradle-wrapper.jar files if present
	files, err = excludeValidatedGradleWrappers(c, files)
	if err != nil {
		return checker.BinaryArtifactData{}, fmt.Errorf("%w", err)
	}

	// No error, return the files.
	return checker.BinaryArtifactData{Files: files}, nil
}

// excludeValidatedGradleWrappers returns the subset of files not confirmed
// to be Action-validated gradle-wrapper.jar files.
func excludeValidatedGradleWrappers(c clients.RepoClient, files []checker.File) ([]checker.File, error) {
	// Check if gradle-wrapper.jar present
	if !fileExists(files, "gradle-wrapper.jar") {
		return files, nil
	}
	// Gradle wrapper JARs present, so check that they are validated
	ok, err := gradleWrapperValidated(c)
	if err != nil {
		return files, fmt.Errorf(
			"failure checking for Gradle wrapper validating Action: %w", err)
	}
	if !ok {
		// Gradle Wrappers not validated
		return files, nil
	}
	// It has been confirmed that latest commit has validated JARs!
	// Remove Gradle wrapper JARs from files.
	for i := range files {
		if filepath.Base(files[i].Path) == "gradle-wrapper.jar" {
			files[i].Type = finding.FileTypeBinaryVerified
		}
	}
	return files, nil
}

var checkBinaryFileReader fileparser.DoWhileTrueOnFileReader = func(path string, reader io.Reader,
	args ...interface{},
) (bool, error) {
	if len(args) != 1 {
		return false, fmt.Errorf(
			"checkBinaryFileReader requires exactly one argument: %w", errInvalidArgLength)
	}
	pfiles, ok := args[0].(*[]checker.File)
	if !ok {
		return false, fmt.Errorf(
			"checkBinaryFileReader requires argument of type *[]checker.File: %w", errInvalidArgType)
	}

	binaryFileTypes := map[string]bool{
		"crx":    true,
		"deb":    true,
		"dex":    true,
		"dey":    true,
		"elf":    true,
		"o":      true,
		"a":      true,
		"so":     true,
		"macho":  true,
		"iso":    true,
		"class":  true,
		"jar":    true,
		"bundle": true,
		"dylib":  true,
		"lib":    true,
		"msi":    true,
		"dll":    true,
		"drv":    true,
		"efi":    true,
		"exe":    true,
		"ocx":    true,
		"pyc":    true,
		"pyo":    true,
		"par":    true,
		"rpm":    true,
		"wasm":   true,
		"whl":    true,
	}

	content, err := io.ReadAll(io.LimitReader(reader, binaryTestLen))
	if err != nil {
		return false, fmt.Errorf("reading file: %w", err)
	}

	var t types.Type
	if len(content) == 0 {
		return true, nil
	}
	if t, err = filetype.Get(content); err != nil {
		return false, sce.WithMessage(sce.ErrScorecardInternal, fmt.Sprintf("filetype.Get:%v", err))
	}

	exists1 := binaryFileTypes[t.Extension]
	if exists1 {
		*pfiles = append(*pfiles, checker.File{
			Path:   path,
			Type:   finding.FileTypeBinary,
			Offset: checker.OffsetDefault,
		})
		return true, nil
	}

	exists2 := binaryFileTypes[strings.ReplaceAll(filepath.Ext(path), ".", "")]
	if !isText(content) && exists2 {
		*pfiles = append(*pfiles, checker.File{
			Path:   path,
			Type:   finding.FileTypeBinary,
			Offset: checker.OffsetDefault,
		})
	}

	return true, nil
}

// determines if the first binaryTestLen bytes are text
//
//	A version of golang.org/x/tools/godoc/util modified to allow carriage returns
//	and utf8.RuneError (0xFFFD), as the file may not be utf8 encoded.
func isText(s []byte) bool {
	const max = binaryTestLen // at least utf8.UTFMax (4)
	if len(s) > max {
		s = s[0:max]
	}
	for i, c := range string(s) {
		if i+utf8.UTFMax > len(s) {
			// last char may be incomplete - ignore
			break
		}
		if c < ' ' && c != '\n' && c != '\t' && c != '\r' {
			// control character - not a text file
			return false
		}
	}
	return true
}

// gradleWrapperValidated checks for the gradle-wrapper-verify action being
// used in a non-failing workflow on the latest commit.
func gradleWrapperValidated(c clients.RepoClient) (bool, error) {
	gradleWrapperValidatingWorkflowFile := ""
	err := fileparser.OnMatchingFileContentDo(c, fileparser.PathMatcher{
		Pattern:       ".github/workflows/*",
		CaseSensitive: false,
	}, checkWorkflowValidatesGradleWrapper, &gradleWrapperValidatingWorkflowFile)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	// no matching files, validation failed
	if gradleWrapperValidatingWorkflowFile == "" {
		return false, nil
	}

	// If validated, check that latest commit has a relevant successful run
	runs, err := c.ListSuccessfulWorkflowRuns(gradleWrapperValidatingWorkflowFile)
	if err != nil {
		// some clients, such as the local file client, don't support this feature
		// claim unvalidated, so that other parts of the check can still be used.
		if errors.Is(err, clients.ErrUnsupportedFeature) {
			return false, nil
		}
		return false, fmt.Errorf("failure listing workflow runs: %w", err)
	}
	commits, err := c.ListCommits()
	if err != nil {
		return false, fmt.Errorf("failure listing commits: %w", err)
	}
	if len(commits) < 1 || len(runs) < 1 {
		return false, nil
	}
	for _, r := range runs {
		if *r.HeadSHA == commits[0].SHA {
			// Commit has corresponding successful run!
			return true, nil
		}
	}

	return false, nil
}

// checkWorkflowValidatesGradleWrapper checks that the current workflow file
// is indeed using the gradle/wrapper-validation-action action, else continues.
func checkWorkflowValidatesGradleWrapper(path string, content []byte, args ...interface{}) (bool, error) {
	validatingWorkflowFile, ok := args[0].(*string)
	if !ok {
		return false, fmt.Errorf("checkWorkflowValidatesGradleWrapper expects arg[0] of type *string: %w", errInvalidArgType)
	}

	action, errs := actionlint.Parse(content)
	if len(errs) > 0 || action == nil {
		// Parse fail, so not this file.
		return true, nil
	}

	for _, j := range action.Jobs {
		for _, s := range j.Steps {
			ea, ok := s.Exec.(*actionlint.ExecAction)
			if !ok {
				continue
			}
			if ea.Uses == nil {
				continue
			}
			if strings.HasPrefix(ea.Uses.Value, "gradle/wrapper-validation-action@") ||
				strings.HasPrefix(ea.Uses.Value, "gradle/actions/wrapper-validation@") {
				// OK! This is it.
				*validatingWorkflowFile = filepath.Base(path)
				return false, nil
			}
		}
	}
	return true, nil
}

// fileExists checks if a file named `name` exists, including within
// subdirectories.
func fileExists(files []checker.File, name string) bool {
	for _, f := range files {
		if filepath.Base(f.Path) == name {
			return true
		}
	}
	return false
}
