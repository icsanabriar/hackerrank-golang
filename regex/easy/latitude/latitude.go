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

// validate function check that the given coordinate comply with specify latitude and longitude.
func validate(w io.Writer, lines []string) {
	re := regexp.MustCompile(`\([+-]?((90(\.0+)?)|((\d|[1-8]\d)(\.\d+)?)), [+-]?((180(\.0+)?)|((\d?\d|1[0-7]\d)(\.\d+)?))\)`)

	for i := 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			_, _ = fmt.Fprintln(w, "Valid")
		} else {
			_, _ = fmt.Fprintln(w, "Invalid")
		}
	}
}
