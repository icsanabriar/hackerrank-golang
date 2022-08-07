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
	"regexp"
	"strings"
)

// ModifyString trim the string, then remove the digits and finally reverse it.
func ModifyString(str string) string {
	result := strings.TrimSpace(str)

	re := regexp.MustCompile(`\d`)
	result = re.ReplaceAllString(result, "")

	reversed := ""
	for _, v := range result {
		reversed = string(v) + reversed
	}

	return reversed
}
