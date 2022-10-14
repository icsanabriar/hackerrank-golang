// Copyright 2022 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
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
	"sort"
	"strings"
)

//extract function to extract the html tags of the given lines.
func extract(w io.Writer, lines []string) {

	re := regexp.MustCompile(`<\s*([a-z][a-z0-9]*)[^>]*>`)
	tags := map[string]bool{}

	for i := 0; i < len(lines); i++ {
		matches := re.FindAllSubmatch([]byte(lines[i]), -1)
		for _, m := range matches {
			tag := string(m[1])
			if _, ok := tags[tag]; !ok {
				tags[tag] = true
			}
		}
	}

	keys := make([]string, 0)

	for k := range tags {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	_, _ = fmt.Fprintln(w, strings.Join(keys, ";"))
}
