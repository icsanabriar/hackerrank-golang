// Copyright 2023 Iván Camilo Sanabria
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

// zigZagSequence generates zig zag array using the given parameter.
func zigZagSequence(arr []int64) []int64 {
	n := len(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	mid := (n+1)/2 - 1
	arr[mid], arr[n-1] = arr[n-1], arr[mid]

	st := mid + 1
	ed := n - 2

	for st <= ed {
		arr[st], arr[ed] = arr[ed], arr[st]
		st++
		ed--
	}

	return arr
}
