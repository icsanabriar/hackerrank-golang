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

import "sort"

// bigSorting return sorted values of the given array using the length of the given values.
func bigSorting(arr []string) []string {

	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) != len(arr[j]) {
			return len(arr[i]) < len(arr[j])
		} else {
			for k := 0; k < len(arr[i]); k++ {
				d1 := arr[i][k] - '0'
				d2 := arr[j][k] - '0'
				if d1 != d2 {
					return d1 < d2
				}
			}
			return false
		}
	})

	return arr
}
