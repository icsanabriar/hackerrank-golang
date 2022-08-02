// Copyright 2021 Iván Camilo Sanabria
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

// MAX Constant to generate the first 15 prime numbers.
const MAX int = 15

// primes Cache to store prime numbers.
var primes = make([]int64, 0)

// isPrime Validates if the given x is prime or not.
func isPrime(x int64) bool {
	for i := int64(2); i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes Generate primes in order to create cache values to improve performance.
func generatePrimes() {

	var prime = int64(2)

	for i := 0; i < MAX; i++ {
		for !isPrime(prime) {
			prime++
		}

		primes = append(primes, prime)
		prime++
	}
}

// primeCount Count the number of primes that are present inside the range 0 to n.
func primeCount(n int64) int32 {

	if len(primes) == 0 {
		generatePrimes()
	}

	result := int32(0)
	marker := int64(1)

	for i := 0; i < len(primes) && marker*primes[i] <= n; i++ {
		result++
		marker = marker * primes[i]
	}

	return result
}
