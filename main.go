// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"

	"github.com/fairwindsops/reckoner/cmd"
)

var (
	// version is set during build
	version = "0.0.0"
	// commit is set during build
	commit = "n/a"
	//go:embed pkg/course/coursev2.schema.json
	courseSchema []byte
)

func main() {
	cmd.Execute(version, commit, courseSchema)
}
