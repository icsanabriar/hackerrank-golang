// Copyright 2020 Iván Camilo Sanabria
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

import "sort"

// cutTheSticks return the number of sticks after each iteration of cutting.
func cutTheSticks(arr []int32) []int32 {

	cache := make(map[int32]int32)

	for i := range arr {
		length := arr[i]
		cache[length]++
	}

	keys := make([]int32, 0, len(cache))

	for k := range cache {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	var result []int32
	left := int32(len(arr))

	for _, k := range keys {
		result = append(result, left)
		left = left - cache[k]
	}

	return result
}
