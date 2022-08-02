// Copyright 2020 Iván Camilo Sanabria
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

// jumpingOnClouds return an integer representing the minimum number of jumps required to get to the end of given
// array c taking into account save jumps.
func jumpingOnClouds(c []int32) int32 {

	i := 0
	jumps := int32(0)

	for i < len(c) {
		if i+2 < len(c) {
			twoSteps := c[i+2]
			if twoSteps == 0 {
				i = i + 2
			} else {
				i = i + 1
			}
		} else {
			i = i + 1
		}
		jumps++
	}

	return jumps - 1
}
