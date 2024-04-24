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

// quickSort return sorted values of the given array using the partition specified on the problem.
func quickSort(arr []int64) []int64 {
	left := make([]int64, 0)
	right := make([]int64, 0)

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[0] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = append(left, arr[0])
	left = append(left, right...)

	return left
}
