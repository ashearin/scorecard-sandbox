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

package evaluation

import (
	"fmt"

	"github.com/ossf/scorecard/v5/checker"
	sce "github.com/ossf/scorecard/v5/errors"
	"github.com/ossf/scorecard/v5/finding"
	"github.com/ossf/scorecard/v5/probes/hasReleaseSBOM"
	"github.com/ossf/scorecard/v5/probes/hasSBOM"
)

// SBOM applies the score policy for the SBOM check.
func SBOM(name string,
	findings []finding.Finding,
	dl checker.DetailLogger,
) checker.CheckResult {
	// We have 4 unique probes, each should have a finding.
	expectedProbes := []string{
		hasSBOM.Probe,
		hasReleaseSBOM.Probe,
	}

	if !finding.UniqueProbesEqual(findings, expectedProbes) {
		e := sce.WithMessage(sce.ErrScorecardInternal, "invalid probe results")
		return checker.CreateRuntimeErrorResult(name, e)
	}

	// Compute the score.
	existsMsg := "SBOM file found in project"
	releaseMsg := "SBOM file found in release artifacts"
	score := 0
	m := make(map[string]bool)
	for i := range findings {
		f := &findings[i]
		switch f.Outcome {
		case finding.OutcomeNotApplicable:
			dl.Info(&checker.LogMessage{
				Type:   finding.FileTypeSource,
				Offset: 1,
				Text:   f.Message,
			})
		case finding.OutcomeTrue:
			switch f.Probe {
			case hasSBOM.Probe:
				dl.Info(&checker.LogMessage{
					Type: finding.FileTypeSource,
					Path: f.Message,
					Text: existsMsg,
				})
				score += scoreProbeOnce(f.Probe, m, 5)
			case hasReleaseSBOM.Probe:
				dl.Info(&checker.LogMessage{
					Type: finding.FileTypeURL,
					Path: f.Message,
					Text: releaseMsg,
				})
				score += scoreProbeOnce(f.Probe, m, 5)
			default:
				e := sce.WithMessage(sce.ErrScorecardInternal, "unknown probe results")
				return checker.CreateRuntimeErrorResult(name, e)
			}
		case finding.OutcomeFalse:
			switch f.Probe {
			case hasSBOM.Probe:
				dl.Warn(&checker.LogMessage{
					Type: finding.FileTypeSource,
					Path: f.Message,
					Text: "SBOM file not found in project",
				})
				existsMsg = f.Message
			case hasReleaseSBOM.Probe:
				dl.Warn(&checker.LogMessage{
					Type: finding.FileTypeURL,
					Path: f.Message,
					Text: "SBOM file not found in release artifacts",
				})
				releaseMsg = f.Message
			}
		default:
			continue // for linting
		}
	}

	_, defined := m[hasSBOM.Probe]
	if !defined {
		return checker.CreateMinScoreResult(name, "SBOM file not detected")
	}

	message := fmt.Sprintf("%s. %s.", existsMsg, releaseMsg)
	return checker.CreateResultWithScore(name, message, score)
}
