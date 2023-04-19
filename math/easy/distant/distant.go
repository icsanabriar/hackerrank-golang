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

import "math"

// findMaxDistance calculates distances using the min and max points over X and Y axes to return the maximum value.
func findMaxDistance(minX, maxX, minY, maxY int32) float64 {

	dx := float64(maxX - minX)
	dy := float64(maxY - minY)

	maxAbs := math.Abs(float64(minX))

	if maxAbs < float64(maxX) {
		maxAbs = float64(maxX)
	}

	mayAbs := math.Abs(float64(minY))

	if mayAbs < float64(maxY) {
		mayAbs = float64(maxY)
	}

	dg := math.Sqrt(math.Pow(maxAbs, 2) + math.Pow(mayAbs, 2))

	if dx > dy {
		if dx > dg {
			return dx
		}
	} else {
		if dy > dg {
			return dy
		}
	}

	return dg
}

// solve calculates the maximum distance of the given set of coordinates.
func solve(coordinates [][]int32) float64 {
	minX := int32(0)
	maxX := int32(0)
	minY := int32(0)
	maxY := int32(0)

	for i := 0; i < len(coordinates); i++ {
		x := coordinates[i][0]
		y := coordinates[i][1]

		if minX > x {
			minX = x
		}

		if minY > y {
			minY = y
		}

		if maxX < x {
			maxX = x
		}

		if maxY < y {
			maxY = y
		}
	}

	return findMaxDistance(minX, maxX, minY, maxY)
}
