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

// flippingMatrix calculates the maximum sum of the upper left quadrant.
func flippingMatrix(matrix [][]int64) int64 {
	n := len(matrix)
	m := len(matrix[0])

	var maxScore int64

	for j := 0; j < n/2; j++ {
		for k := 0; k < m/2; k++ {
			a := matrix[j][k]
			b := matrix[j][m-1-k]
			c := matrix[n-1-j][k]
			d := matrix[n-1-j][m-1-k]
			currentScore := a

			if b > currentScore {
				currentScore = b
			}
			if c > currentScore {
				currentScore = c
			}
			if d > currentScore {
				currentScore = d
			}

			maxScore += currentScore
		}
	}

	return maxScore
}
