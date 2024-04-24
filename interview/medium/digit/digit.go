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

import (
	"strconv"
)

// superDigit calculate the sum of the given digits embedded on the string.
func superDigit(s string) int {
	sum := 0
	for _, r := range s {
		sum += int(r - '0')
	}

	if sum < 10 {
		return sum
	}

	return superDigit(strconv.Itoa(sum))
}
