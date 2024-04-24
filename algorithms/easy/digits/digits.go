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

// findDigits return an integer representing the number of digits of d that are divisor of given n.
func findDigits(n int64) int {
	divisorDigit := false
	counter := 0
	originalN := n
	previousDigit := make(map[int64]bool)

	for n > 0 {
		currentDigit := n % 10
		if currentDigit != 0 {
			if val, ok := previousDigit[currentDigit]; !ok {
				divisorDigit = (originalN % currentDigit) == 0
				previousDigit[currentDigit] = divisorDigit
				if divisorDigit {
					counter++
				}
			} else if val {
				counter++
			}
		}
		n /= 10
	}

	return counter
}
