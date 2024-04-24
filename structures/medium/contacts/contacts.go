// Copyright 2020 Iván Camilo Sanabria
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

// contacts count the partial results on cache map given on queries, which is a list of commands add or find.
func contacts(queries [][]string) []int32 {
	// Build cache on initial state.
	cache := make(map[string]int32)
	result := make([]int32, 0)

	// Iterate over queries.
	for _, row := range queries {
		// Add possible partials to cache.
		if row[0] == "add" {
			for i := range row[1] {
				currentString := row[1][0 : i+1]
				if val, ok := cache[currentString]; ok {
					cache[currentString] = val + 1
				} else {
					cache[currentString] = 1
				}
			}
		}
		// Add values to result when there is a partial
		if row[0] == "find" {
			result = append(result, cache[row[1]])
		}
	}

	return result
}
