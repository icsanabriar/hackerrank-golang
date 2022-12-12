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

import (
	"math"
	"sort"
)

// event to store in cache to find cloud to remove.
type event struct {
	t   int
	x   int64
	p   int64
	idx int64
}

// maximumPeople estimates the maximum number of people that will be in a sunny town after removing exactly one cloud.
func maximumPeople(p []int64, x []int64, y []int64, r []int64) int64 {
	free := int64(0)
	events := generateCache(p, x, y, r)

	sum := make(map[int64]int64, 0)
	clouds := make(map[int64]bool, 0)

	for _, e := range events {
		if e.t == 0 {
			clouds[e.idx] = true
		} else if e.t == 1 {
			if len(clouds) == 0 {
				free += e.p
			} else {
				if len(clouds) == 1 {
					for k := range clouds {
						sum[k] = sum[k] + e.p
					}
				}
			}
		} else {
			delete(clouds, e.idx)
		}
	}

	return free + findMax(sum)
}

// generateCache build the event cache and sort based on x and t values.
func generateCache(p []int64, x []int64, y []int64, r []int64) []event {
	events := make([]event, 0)

	for i := 0; i < len(p); i++ {
		e := event{
			t:   1,
			x:   x[i],
			p:   p[i],
			idx: int64(i),
		}

		events = append(events, e)
	}

	for i := 0; i < len(y); i++ {
		e0 := event{
			t:   0,
			x:   y[i] - r[i],
			p:   -1,
			idx: int64(i),
		}

		events = append(events, e0)

		e2 := event{
			t:   2,
			x:   y[i] + r[i],
			p:   -1,
			idx: int64(i),
		}

		events = append(events, e2)
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].x != events[j].x {
			return events[i].x < events[j].x
		}
		return events[i].t < events[j].t
	})

	return events
}

// findMax find the maximum value form the given map.
func findMax(sum map[int64]int64) int64 {
	dx := int64(0)

	for _, v := range sum {
		dx = int64(math.Max(float64(dx), float64(v)))
	}
	return dx
}
