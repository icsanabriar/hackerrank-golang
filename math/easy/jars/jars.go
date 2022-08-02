// Copyright 2020 Iván Camilo Sanabria
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

// estimate calculates the average of candies in a set of n jars by executing adding operations
// given index a,b and number of candies to add.
func estimate(n int64, operations [][]int64) int64 {
	result := int64(0)

	for i := range operations {
		a := operations[i][0]
		b := operations[i][1]
		k := operations[i][2]

		result += (b - a + 1) * k
	}

	return result / n
}
