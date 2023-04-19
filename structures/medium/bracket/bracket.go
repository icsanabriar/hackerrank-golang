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

// isBalanced validates that the given string has complete inner pairs of brackets.
func isBalanced(s string) string {
	cache := make([]string, 0)

	for i := range s {
		char := string(s[i])
		if char == "(" || char == "[" || char == "{" {
			cache = append(cache, char)
		} else if len(cache) > 0 {
			last := len(cache) - 1
			if (char == ")" && cache[last] != "(") || (char == "]" && cache[last] != "[") || (char == "}" && cache[last] != "{") {
				return "NO"
			}
			cache = cache[:last]
		} else {
			return "NO"
		}
	}

	if len(cache) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
