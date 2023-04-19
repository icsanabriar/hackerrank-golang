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

// swap elements from the given array using the given indexes.
func swap(arr []int32, i int32, j int32) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// minimumSwaps count the min swaps required to sort the given array.
func minimumSwaps(arr []int32) int32 {
	cache := map[int32]int32{}

	for i := int32(0); i < int32(len(arr)); i++ {
		cache[arr[i]] = i
	}

	counter := int32(0)

	for j := int32(0); j < int32(len(arr)); j++ {
		if arr[j] != j+1 {
			index := cache[j+1]
			cache[arr[j]] = index
			cache[j+1] = j
			swap(arr, j, index)
			counter++
		}
	}

	return counter
}
