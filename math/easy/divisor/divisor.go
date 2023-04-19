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

// sumDigits return an integer representing the sum of the given n.
func sumDigits(n int32) int32 {
	sum := int32(0)

	for n > 0 {
		digit := n % 10
		sum = sum + digit
		n = n / 10
	}

	return sum
}

// findDivisor return an integer representing the best sum of digits from a divisor of given n.
func findDivisor(n int32) int32 {

	maximum := int32(0)
	result := int32(0)

	for i := int32(1); i*i <= n; i++ {

		// Validate index is a divisor of n.
		if n%i == 0 {

			di := sumDigits(i)

			if di > maximum || (di == maximum && result > i) {
				result = i
				maximum = di
			}

			dii := sumDigits(n / i)

			if dii > maximum || (dii == maximum && result > n/i) {
				result = n / i
				maximum = dii
			}
		}
	}

	return result
}
