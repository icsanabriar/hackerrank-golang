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

import "math"

// diagonalDifference calculates the difference between the diagonals of the given matrix.
func diagonalDifference(arr [][]int32) int32 {

	var sumPrimaryDiagonal int32
	var sumSecondaryDiagonal int32

	n := len(arr)

	for i := 0; i < n; i++ {
		sumPrimaryDiagonal += arr[i][i]
		sumSecondaryDiagonal += arr[i][n-1-i]
	}

	return int32(math.Abs(float64(sumPrimaryDiagonal - sumSecondaryDiagonal)))
}
