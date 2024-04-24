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

// minFactor calculates the given n could be decomposed.
func minFactor(n int64) int64 {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

// gcd calculate the great common divisor of given a,b.
func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// easyGcd Finds the maximum number that could be part of the arr given the k boundary.
func easyGcd(k int64, arr []int64) int64 {
	g := arr[0]
	for i := 1; i < len(arr); i++ {
		g = gcd(g, arr[i])
	}

	minfact := minFactor(g)

	return k / minfact * minfact
}
