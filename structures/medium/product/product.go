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

// makeRange build a range between a and b including them. a > b otherwise the function panic.
func makeRange(a, b int32) []int32 {

	arr := make([]int32, b-a)

	for i := range arr {
		arr[i] = a + int32(i)
	}

	return arr
}

// reverse exchange the order of the elements of the given arr in reverse mode.
// Example: [1,2,3,4] -> [4,3,2,1].
func reverse(arr []int32) []int32 {

	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}

	return arr
}

// buildIndexes calculates the near element that is higher than arr[i]. Depending on the value of indexes
// it goes to the right or left direction.
func buildIndexes(arr, indexes []int32) []int32 {

	nearest := make([]int32, len(arr))
	temp := make([]int32, 0)

	for e := range indexes {

		for len(temp) > 0 && arr[indexes[e]] >= arr[temp[len(temp)-1]] {
			temp = temp[:len(temp)-1]
		}

		nearest[indexes[e]] = int32(0)

		if len(temp) > 0 {
			nearest[indexes[e]] = temp[len(temp)-1] + 1
		}

		temp = append(temp, indexes[e])
	}

	return nearest
}

// solve calculates the maximum index product of the given array.
func solve(arr []int32) int64 {

	indexes := makeRange(0, int32(len(arr)))
	leftIndex := buildIndexes(arr, indexes)
	rightIndex := buildIndexes(arr, reverse(indexes))

	maximum := int64(0)

	for i := 0; i < len(arr); i++ {

		// Casting is required before product for large values.
		product := int64(leftIndex[i]) * int64(rightIndex[i])

		if product > maximum {
			maximum = product
		}
	}

	return maximum
}
