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

// sockMerchant count the number of pairs given in the array.
func sockMerchant(n int64, arr []int64) int64 {
	cache := map[int64]int{}
	pairs := int64(0)

	for i := int64(0); i < n; i++ {
		if _, ok := cache[arr[i]]; ok {
			delete(cache, arr[i])
			pairs++
		} else {
			cache[arr[i]] = 1
		}
	}

	return pairs
}
