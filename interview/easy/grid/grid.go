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

import "sort"

// gridChallenge validates that the given grid is order en each column after sorting the rows.
func gridChallenge(grid []string) string {
	for row := range grid {
		b := []byte(grid[row])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		grid[row] = string(b)
	}

	for i := range grid[0] {
		for j := 1; j < len(grid); j++ {
			if grid[j][i] < grid[j-1][i] {
				return "NO"
			}
		}
	}

	return "YES"
}
