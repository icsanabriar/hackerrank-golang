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

// truckTour calculates the starting pump the truck could take a full tour.
func truckTour(pumps [][]int32) int32 {

	var start, fuel, distance int32

	for i := 0; i < len(pumps); i++ {
		fuel += pumps[i][0] - pumps[i][1]
		distance += pumps[i][0] - pumps[i][1]

		if fuel < 0 {
			start = int32(i) + 1
			fuel = 0
		}
	}

	if fuel == 0 {
		return -1
	}

	return start
}
