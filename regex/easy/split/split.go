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

// classify function extract information of the given phone numbers into specify format.
func classify(w io.Writer, a []string) {

	re := regexp.MustCompile(`(\d{1,3})[- ](\d{1,3})[- ](\d{4,10})`)

	for tc := 0; tc < len(a); tc++ {
		match := re.FindStringSubmatch(a[tc])
		code := match[1]
		area := match[2]
		number := match[3]
		_, _ = fmt.Fprintf(w, "CountryCode=%v,LocalAreaCode=%v,Number=%v\n", code, area, number)
	}
}
