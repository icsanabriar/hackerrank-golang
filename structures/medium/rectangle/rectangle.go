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

// largestRectangle return an integer representing the largest rectangle that can be formed within the bounds of
// consecutive buildings given the heights on array h.
func largestRectangle(h []int32) int64 {

	largest := int64(0)
	var nearest int32

	for i, hi := range h {
		nearest = 1

		// Go to the right and check nearest.
		for j := i + 1; j < len(h); j++ {
			if hi > h[j] {
				break
			} else {
				nearest += 1
			}
		}

		// Go to the left and check nearest.
		for j := i - 1; j >= 0; j-- {
			if hi > h[j] {
				break
			} else {
				nearest += 1
			}
		}

		area := int64(hi * nearest)

		if area > largest {
			largest = area
		}
	}

	return largest
}
