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
	"reflect"
	"sort"
)

// isEqual validates that the given array are equal when those are sorted.
func isEqual(a, b []int64) bool {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	return reflect.DeepEqual(a, b)
}

// organizingContainers find if the given container are possible to sort or not.
func organizingContainers(container [][]int64, size int64) string {
	rowSum := make([]int64, size)
	for i := int64(0); i < size; i++ {
		for j := int64(0); j < size; j++ {
			rowSum[i] += container[i][j]
		}
	}

	colSum := make([]int64, size)
	for i := int64(0); i < size; i++ {
		for j := int64(0); j < size; j++ {
			colSum[i] += container[j][i]
		}
	}

	if isEqual(rowSum, colSum) {
		return "Possible"
	}

	return "Impossible"
}
