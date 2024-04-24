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

import "sort"

// pickingNumbers return the maximum number of integers you can select from the array that the absolute difference
// between any two of the chosen integers is less than or equal to 1.
func pickingNumbers(a []int64) int64 {
	m := make(map[int64]int64)

	// Memorize array in order and count the number of each int.
	for i := range a {
		m[a[i]]++
	}

	// Basic case when there is one element.
	if len(m) == 1 {
		return int64(len(a))
	}

	keys := make([]int64, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys of the map to calculate maximum length.
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return maxLength(m, keys)
}

// maxLength Iterate over map looking for integers which absolute difference is less or equal to 1.
func maxLength(m map[int64]int64, keys []int64) int64 {
	maximumLength := m[keys[0]]

	for i := 0; i < (len(keys) - 1); i++ {
		if keys[i+1]-keys[i] <= 1 {
			currentLength := m[keys[i+1]] + m[keys[i]]

			if currentLength > maximumLength {
				maximumLength = currentLength
			}
		} else {
			currentLength := m[keys[i+1]]

			if currentLength > maximumLength {
				maximumLength = currentLength
			}
		}
	}

	return maximumLength
}
