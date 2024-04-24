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

// matchingStrings count the given queries on the string arrays. The query and string should be the same.
func matchingStrings(strings []string, queries []string) []int32 {
	cache := make(map[string]int32)
	// Count the number of strings and store in a map.
	for _, e := range strings {
		if val, ok := cache[e]; ok {
			cache[e] = val + 1
		} else {
			cache[e] = 1
		}
	}

	result := make([]int32, 0)
	// Find matches of queries.
	for _, e := range queries {
		result = append(result, cache[e])
	}

	return result
}
