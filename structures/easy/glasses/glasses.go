// Copyright 2020 Iván Camilo Sanabria
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

// hourglassSum calculate the maximum hourglass sum in the given arr.
func hourglassSum(arr [][]int64) int64 {
	maximum := int64(-64)
	var i, j int
	var sum int64

	// Iterate over rows of arr.
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			// First row of the glass.
			sum = arr[i][j]
			sum += arr[i][j+1]
			sum += arr[i][j+2]
			// Second row of the glass.
			sum += arr[i+1][j+1]
			// Third row of the glass.
			sum += arr[i+2][j]
			sum += arr[i+2][j+1]
			sum += arr[i+2][j+2]
			// Check maximum of the glass.
			if maximum < sum {
				maximum = sum
			}
		}
	}

	return maximum
}
