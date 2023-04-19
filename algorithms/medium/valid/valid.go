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

// min retrieves the minimum value of the given values.
func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

// max retrieves the maximum value of the given values.
func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

// isValid validates that the given string contains the same amount of characters, or maximum 1 delta.
func isValid(s string) string {
	cache := make(map[string]int32)

	for i := range s {
		key := string(s[i])
		if value, ok := cache[key]; ok {
			cache[key] = value + 1
		} else {
			cache[key] = 1
		}
	}

	reduce := make(map[int32]int32)
	minKey := int32(2147483647)
	maxKey := int32(0)

	for _, v := range cache {
		if value, ok := reduce[v]; ok {
			reduce[v] = value + 1
		} else {
			reduce[v] = 1
		}
		minKey = min(minKey, v)
		maxKey = max(maxKey, v)
	}

	length := len(reduce)

	if length == 1 {
		return "YES"
	}

	if (maxKey-minKey == 1) && (reduce[maxKey] == 1 || reduce[minKey] == 1) {
		return "YES"
	}

	if length == 2 && minKey == 1 && reduce[minKey] == 1 {
		return "YES"
	}

	return "NO"
}
