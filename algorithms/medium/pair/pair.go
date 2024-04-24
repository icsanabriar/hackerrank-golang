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

// pairs search the number of pair in the given array with the same difference as k.
func pairs(k int64, arr []int64) int64 {
	cache := make(map[int64]bool)
	keys := make([]int64, 0, len(cache))

	for i := range arr {
		if _, ok := cache[arr[i]]; !ok {
			cache[arr[i]] = true
			keys = append(keys, arr[i])
		}
	}

	counter := int64(0)

	for i := range keys {
		key := keys[i]
		desiredKey := key + k
		if value, ok := cache[desiredKey]; ok && value {
			cache[desiredKey] = false
			counter++
		}
	}

	return counter
}
