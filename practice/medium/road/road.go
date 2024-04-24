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

// roadsAndLibraries estimates the minimum cost to build libraries for all cities.
func roadsAndLibraries(n int64, lib int64, road int64, cities [][]int64) int64 {
	if road >= lib {
		return n * lib
	}

	cost := int64(0)
	cacheCity := make(map[int64]int64)
	cacheGroup := make(map[int64][]int64)

	for i := int64(1); i <= n; i++ {
		cacheCity[i] = i
		cacheGroup[i] = []int64{i}
	}

	for _, r := range cities {
		c1 := r[0]
		c2 := r[1]
		g1 := cacheCity[c1]
		g2 := cacheCity[c2]

		if g1 != g2 {
			cacheCity[c2] = g1
			for _, c := range cacheGroup[g2] {
				cacheCity[c] = g1
			}
			cacheGroup[g1] = append(cacheGroup[g1], cacheGroup[g2]...)
			delete(cacheGroup, g2)
			cost += road
		}
	}

	return cost + (int64(len(cacheGroup)) * lib)
}
