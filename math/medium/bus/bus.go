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

import "sort"

// sum add all the values of the array and return the total.
func sum(a []int32) int32 {
	total := int32(0)

	for i := range a {
		total = total + a[i]
	}

	return total
}

// validate checks if the given size fits for all elements in array a.
func validate(a []int32, size int32) bool {
	remain := size

	for i := range a {
		if remain == a[i] {
			remain = size
		} else if remain < a[i] {
			return false
		} else {
			remain = remain - a[i]
		}
	}

	return remain == size
}

// solve return the size of the bus to fit all the given groups in multiple rides.
func solve(a []int32) []int32 {
	total := sum(a)
	sizes := make([]int32, 0)

	for i := int32(1); i*i <= total; i++ {
		if total%i == 0 {
			if validate(a, i) {
				sizes = append(sizes, i)
			}
			if i*i != total {
				if validate(a, total/i) {
					sizes = append(sizes, total/i)
				}
			}
		}
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] < sizes[j] })

	return sizes
}
