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

// findWay retrieves a way to exchange the given remain.
func findWay(c []int64, remain int64, index int, cache [][]int64) int64 {
	if remain == 0 {
		return 1
	}

	if index == 0 {
		if remain%c[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	if cache[remain][index] >= 0 {
		return cache[remain][index]
	}

	result := int64(0)

	for i := int64(0); (i * c[index]) <= remain; i++ {
		result += findWay(c, remain-i*c[index], index-1, cache)
	}

	cache[remain][index] = result

	return result
}

// getWays find the number of possible change based on the given amount n and the coin designation c.
func getWays(n int64, c []int64) int64 {

	cache := make([][]int64, n+1)

	for i := range cache {
		cache[i] = make([]int64, len(c))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	return findWay(c, n, len(c)-1, cache)
}
