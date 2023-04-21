// Copyright 2023 Iván Camilo Sanabria
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

// caesarCipher encrypt the given string using the k factor.
func caesarCipher(s string, k int32) string {

	text := ""

	for _, char := range s {

		if char >= 'a' && char <= 'z' {
			char = (char-'a'+k)%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			char = (char-'A'+k)%26 + 'A'
		}

		text += string(char)
	}

	return text
}
