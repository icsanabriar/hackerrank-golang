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

// road present in gridland.
type road struct {
	c1 int32
	c2 int32
}

// roadSpace estimates how many space takes the given array on gridland.
func roadSpace(roads []road) int64 {

	sort.Slice(roads, func(i int, j int) bool {
		if roads[i].c1 == roads[j].c1 {
			return roads[i].c2 < roads[j].c2
		} else {
			return roads[i].c1 < roads[j].c1
		}
	})

	spaces := int64(0)
	max := int32(0)

	for i := 0; i < len(roads); i++ {
		r := roads[i]
		if i == 0 {
			spaces += int64(r.c2 - r.c1 + 1)
			max = r.c2
		} else {
			if r.c2 > max {
				if r.c1 <= max {
					spaces += int64(r.c2 - max)
				} else {
					spaces += int64(r.c2 - r.c1 + 1)
				}
				max = r.c2
			}
		}
	}

	return spaces
}

// gridlandMetro calculates the number of cells where lampposts can be installed.
func gridlandMetro(n int32, m int32, track [][]int32) int64 {

	cache := make(map[int32][]road)

	for _, t := range track {
		row := t[0]
		r := road{t[1], t[2]}

		if v, ok := cache[row]; ok {
			cache[row] = append(v, r)
		} else {
			cache[row] = []road{r}
		}
	}

	cells := int64(n) * int64(m)

	for _, r := range cache {
		cells = cells - roadSpace(r)
	}

	return cells
}
