// Copyright 2021 Iván Camilo Sanabria
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

// movingTiles calculates the area of intersection square between two tiles.
func movingTiles(l int32, s1 int32, s2 int32, queries []int64) []float64 {

	results := make([]float64, 0)

	for i := range queries {
		q := (float64(l)*math.Sqrt2 - math.Sqrt(float64(queries[i]))*math.Sqrt2) / math.Abs(float64(s1)-float64(s2))
		results = append(results, q)
	}

	return results
}
