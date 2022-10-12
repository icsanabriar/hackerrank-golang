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
	"strings"
)

//count function to count the number the terms appear on the text taking into account uk and usa spelling.
func count(w io.Writer, text string, terms []string) {

	for t := 0; t < len(terms); t++ {
		term := terms[t]

		re := regexp.MustCompile("\\b(" + term + "|" + strings.ReplaceAll(term, "our", "or") + ")\\b")
		_, _ = fmt.Fprintf(w, "%v\n", len(re.FindAll([]byte(text), -1)))
	}
}
