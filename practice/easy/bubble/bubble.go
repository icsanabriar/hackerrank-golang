// Copyright 2022 Iván Camilo Sanabria
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

import "fmt"

// swap elements from the given array using the given indexes.
func swap(a []int32, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

// countSwaps implements the count of swaps to sort the given array.
func countSwaps(a []int32) string {
	counter := 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				swap(a, j, j+1)
				counter++
			}
		}
	}

	result := fmt.Sprintf("Array is sorted in %v swaps.\n", counter)
	result = result + fmt.Sprintf("First Element: %v\n", a[0])
	result = result + fmt.Sprintf("Last Element: %v\n", a[len(a)-1])

	return result
}
