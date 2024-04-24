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

import (
	"math"
	"sort"
)

// closestNumbers find the closest pairs of the given array.
func closestNumbers(arr []int64) []int64 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	minimum := int64(math.MaxInt32)
	indexes := make([]int64, 0)

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minimum {
			minimum = diff
			indexes = make([]int64, 0)
			indexes = append(indexes, int64(i))
		} else if diff == minimum {
			indexes = append(indexes, int64(i))
		}
	}

	results := make([]int64, 0)
	for i := 0; i < len(indexes); i++ {
		results = append(results, arr[indexes[i]], arr[indexes[i]+1])
	}

	return results
}
