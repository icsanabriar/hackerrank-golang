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

// solve count the number of ascending triplets present in the given array and ratio.
func countTriplets(arr []int64, r int64) int64 {

	countMap := make(map[int64]int64)
	pairsMap := make(map[int64]int64)

	var count int64

	for _, num := range arr {

		if pairsMap[num] > 0 {
			count += pairsMap[num]
		}

		if countMap[num] > 0 {
			pairsMap[num*r] += countMap[num]
		}

		countMap[num*r]++
	}

	return count
}
