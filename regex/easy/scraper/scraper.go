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

//scrape function that extract questions of the provided text.
func scrape(w io.Writer, text string) {

	re := regexp.MustCompile(`(?s)<h3><a .*?/questions/(\d{5})/.*?class="question-hyperlink">(.+?)</a>.*?class="relativetime">(.+?)</span>`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		_, _ = fmt.Fprintf(w, "%s;%s;%s\n", match[1], match[2], match[3])
	}
}
