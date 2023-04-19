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

// candies calculates the minimum number of candies to buy.
func candies(n int32, arr []int32) int64 {

	c := make([]int32, n)
	c[0] = 1

	for i := 1; i < int(n); i++ {
		c[i] = 1
		if arr[i] > arr[i-1] {
			c[i] = c[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] && c[i] <= c[i+1] {
			c[i] = c[i+1] + 1
		}
	}

	total := int64(0)

	for i := 0; i < int(n); i++ {
		total += int64(c[i])
	}

	return total
}
