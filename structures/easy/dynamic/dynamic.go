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

//dynamicArray uses the given queries to complete arrays using the given n.
func dynamicArray(n int32, queries [][]int32) []int32 {
	last := int32(0)
	results := make([]int32, 0)
	dynamic := make([][]int32, n)

	for _, item := range queries {
		operation := item[0]
		x := item[1]
		y := item[2]

		index := (x ^ last) % n

		if 1 == operation {
			array := dynamic[index]
			dynamic[index] = append(array, y)
		} else if 2 == operation {
			array := dynamic[index]
			index = y % int32(len(array))
			last = array[index]
			results = append(results, last)
		}
	}

	return results
}
