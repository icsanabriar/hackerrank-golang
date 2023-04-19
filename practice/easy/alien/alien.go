// Copyright 2022 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"io"
	"regexp"
)

// validate function check that the given lines contains a valid alien username.
func validate(w io.Writer, lines []string) {

	re := regexp.MustCompile(`^[_.][0-9]+[a-zA-Z]*_?$`)

	for i := 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			_, _ = fmt.Fprintln(w, "VALID")
		} else {
			_, _ = fmt.Fprintln(w, "INVALID")
		}
	}
}
