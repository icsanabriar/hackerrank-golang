// Copyright 2022 Iván Camilo Sanabria
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

import "strconv"

// solve find that the given dates are the same based on the number and base used for each one.
func solve(dates [][]int32) int64 {

	cache := make(map[int64]int64, 0)

	for _, date := range dates {
		month := int(date[0])
		day := int(date[1])

		dayString := strconv.FormatInt(int64(day), 10)
		monthBase, err := strconv.ParseInt(dayString, month, 64)

		if err == nil {
			cache[monthBase]++
		}
	}

	max := int64(0)

	for _, v := range cache {
		max += v * (v - 1) / 2
	}

	return max
}
