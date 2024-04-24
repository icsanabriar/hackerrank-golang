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

// hurdleRace return an integer representing the minimum number of doses required to jump all the hurdle.
func hurdleRace(k int64, height []int64) int64 {
	maximum := height[0]
	for _, value := range height {
		if value > maximum {
			maximum = value
		}
	}

	if (maximum - k) < 0 {
		return 0
	}

	return maximum - k
}
