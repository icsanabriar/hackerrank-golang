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

// miniMaxSum find the min and maximum sum of the given array.
func miniMaxSum(arr []int32) (int64, int64) {
	var minSum, maxSum, currentSum int64

	minSum = int64(arr[0])
	maxSum = int64(arr[0])

	for _, val := range arr {
		currentSum += int64(val)

		if int64(val) < minSum {
			minSum = int64(val)
		}
		if int64(val) > maxSum {
			maxSum = int64(val)
		}
	}

	return currentSum - maxSum, currentSum - minSum
}
