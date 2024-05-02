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

package clients

import (
	"regexp"
	"time"
)

// SBOM represents a customized struct for SBOM used by clients.
type SBOM struct {
	ID            string
	Name          string
	Origin        string
	URL           string
	Created       time.Time
	Path          string
	Tool          string
	Schema        string
	SchemaVersion string
}

var ReSBOMFile = regexp.MustCompile(
	`(?i).+\.(cdx.json|cdx.xml|spdx|spdx.json|spdx.xml|spdx.y[a?]ml|spdx.rdf|spdx.rdf.xm)`,
)
