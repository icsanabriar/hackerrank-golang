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

import "math"

// gcd calculate the great common divisor of given a,b.
func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// diffGcd calculate gcd of the absolute difference of 2 elements in position i and i-1.
func diffGcd(a []int64) int64 {
	result := int64(math.Abs(float64(a[1] - a[0])))

	for i := 2; i < len(a); i++ {
		result = gcd(result, int64(math.Abs(float64(a[i]-a[i-1]))))
	}

	return result
}

// normalize finds the common salaries for the given array, taking into account the given queries
// that represents raises for all employees of a before normalization is executed.
func normalize(a []int64, queries []int64) []int64 {
	result := make([]int64, 0)
	diff := diffGcd(a)

	for i := 0; i < len(queries); i++ {
		k := queries[i]
		result = append(result, gcd(diff, a[0]+k))
	}

	return result
}
