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

// lonelyInteger retrieves the integer of the array that is frequent one time.
func lonelyInteger(arr []int64) int64 {
	cache := make(map[int64]int64)
	for _, num := range arr {
		cache[num]++
	}

	for key, value := range cache {
		if value == 1 {
			return key
		}
	}

	return 0
}
