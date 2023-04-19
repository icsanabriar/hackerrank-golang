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

// fibonacci return the sequence of Fibonacci with the given length n. Cache is map used to memorize the sequence
// that is already estimated on previous loops of recursion in order to increase performance of the function.
func fibonacci(n int64, cache []int64) int64 {

	if n == 0 || n == 1 {
		return n
	}

	if cache[n] == -1 {
		cache[n] = fibonacci(n-1, cache) + fibonacci(n-2, cache)
	}

	return cache[n]
}

// solveIsFibo defines if the given n is a number part of Fibonacci sequence or not.
func solveIsFibo(n int64, cache []int64) string {

	// Memorize first 99 fibonacci numbers.
	_ = fibonacci(99, cache)

	// Iterate in order to optimize search.
	for _, fib := range cache {

		// Current Fibonacci value is higher than given number than break.
		if n < fib {
			break
		}

		// Is Fibonacci return value.
		if n == fib {
			return "IsFibo"
		}
	}

	return "IsNotFibo"
}
