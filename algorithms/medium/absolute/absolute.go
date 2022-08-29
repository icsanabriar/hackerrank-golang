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

// absolutePermutation retrieves the array with n elements which difference between index and element at the index is k.
// In case is not possible it returns an array with -1.
func absolutePermutation(n int32, k int32) []int32 {

	result := make([]int32, 0)

	if k > 0 && n%(2*k) != 0 {
		result = append(result, -1)
		return result
	}

	a := make([]int32, 0)
	cache := make([]bool, 0)

	for i := int32(1); i <= n; i++ {
		a = append(a, i)
		cache = append(cache, false)
		result = append(result, 0)
	}

	for i := int32(0); i < int32(len(a)); i++ {
		if !cache[i] {
			result[i] = a[i+k]
			result[i+k] = a[i]
			cache[i] = true
			cache[i+k] = true
		}
	}

	return result
}
