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

// minimum retrieves the min value of the given values.
func minimum(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// maximum retrieves the max value of the given values.
func maximum(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

// nonDivisibleSubset find the array S that is not divisible by k.
func nonDivisibleSubset(k int64, s []int64) int64 {
	modules := make([]int64, k)
	for i := range s {
		modules[s[i]%k]++
	}

	result := int64(0)
	for i := int64(0); i*2 <= k; i++ {
		opposite := (k - i) % k
		if i == opposite {
			result += minimum(modules[i], 1)
		} else {
			result += maximum(modules[i], modules[opposite])
		}
	}

	return result
}
