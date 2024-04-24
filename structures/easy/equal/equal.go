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

// Sum the values of the given array.
func sumHeights(h []int64) int64 {
	total := int64(0)
	for i := 0; i < len(h); i++ {
		total += h[i]
	}

	return total
}

// Poll element from array.
func pollElement(h1 *[]int64, h2 *[]int64, h3 *[]int64, index int) int64 {
	element := int64(0)

	switch {
	case index == 0 && len(*h1) > 0:
		element = (*h1)[0]
		*h1 = (*h1)[1:]
	case index == 1 && len(*h2) > 0:
		element = (*h2)[0]
		*h2 = (*h2)[1:]
	default:
		element = (*h3)[0]
		*h3 = (*h3)[1:]
	}

	return element
}

// equalStacks calculate the maximum height the three given stacks could be equal.
func equalStacks(h1 []int64, h2 []int64, h3 []int64) int64 {
	cache := make([]int64, 0)

	cache = append(cache, sumHeights(h1))
	cache = append(cache, sumHeights(h2))
	cache = append(cache, sumHeights(h3))

	for !(cache[0] == cache[1] && cache[1] == cache[2]) {
		maxHeightIndex := 0
		for i := 1; i < len(cache); i++ {
			if cache[i] > cache[maxHeightIndex] {
				maxHeightIndex = i
			}
		}
		cache[maxHeightIndex] -= pollElement(&h1, &h2, &h3, maxHeightIndex)
	}

	return cache[0]
}
