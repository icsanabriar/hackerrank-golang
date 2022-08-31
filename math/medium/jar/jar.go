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

// gcd calculate the great common divisor of given a,b.
func gcd(a int32, b int32) int32 {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// max retrieves the maximum value of the given values.
func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

// solve validates that the given amount c gallons could be achieve by the a and b gallon jugs.
func solve(a int32, b int32, c int32) string {
	divisor := gcd(a, b)

	if c <= max(a, b) && c%divisor == 0 {
		return "YES"
	}

	return "NO"
}
