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
)

//find function to validate if the given strings starts, ends or both with hackerrank term.
func find(w io.Writer, a []string) {

	reb := regexp.MustCompile(`^hackerrank$`)
	res := regexp.MustCompile(`^hackerrank`)
	ree := regexp.MustCompile(`hackerrank$`)

	for i := 0; i < len(a); i++ {
		if reb.MatchString(a[i]) {
			_, _ = fmt.Fprintf(w, "0\n")
		} else if res.MatchString(a[i]) {
			_, _ = fmt.Fprintf(w, "1\n")
		} else if ree.MatchString(a[i]) {
			_, _ = fmt.Fprintf(w, "2\n")
		} else {
			_, _ = fmt.Fprintf(w, "-1\n")
		}
	}
}
