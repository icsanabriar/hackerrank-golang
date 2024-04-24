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

// gcd calculate the great common divisor of given a,b.
func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// maximum retrieves the max value of the given values.
func maximum(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

// solve validates that the given amount c gallons could be achieved by the a and b gallon jugs.
func solve(a int64, b int64, c int64) string {
	divisor := gcd(a, b)

	if c <= maximum(a, b) && c%divisor == 0 {
		return "YES"
	}

	return "NO"
}
