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

import (
	"math/big"
	"sort"
)

// prime calculates the i-th prime using a starting int.
func prime(previous int64) int64 {
	for i := previous + 1; ; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			return i
		}
	}
}

// waiter separates the given array of plates (number) into Ai Bi piles doing q iterations.
// Returns an array of int with the order of the piles B0, B1, ... Bi, Ai taking into account
// the TOP element which is the last one of the pile.
func waiter(number []int64, q int64) []int64 {
	a := make(map[int][]int64)
	b := make(map[int][]int64)
	p := int64(2)

	// Build piles from ai.
	for i := 1; i <= int(q); i++ {
		for j := len(number) - 1; j >= 0; j-- {
			e := number[j]
			if e%p == 0 {
				appendValueToArray(b, i, e)
			} else {
				appendValueToArray(a, i, e)
			}
		}
		number = a[i]
		delete(a, i)
		p = prime(p)
	}

	return buildPiles(number, b)
}

// appendValueToArray append the given value to the map using the given key.
func appendValueToArray(m map[int][]int64, k int, v int64) {
	if val, ok := m[k]; ok {
		m[k] = append(val, v)
	} else {
		m[k] = []int64{v}
	}
}

// buildPiles append the results of the piles using the given number and indexes.
func buildPiles(number []int64, b map[int][]int64) []int64 {
	result := make([]int64, 0)
	// Sort keys of b.
	keys := make([]int, 0, len(b))
	for k := range b {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	// Append b values.
	for _, k := range keys {
		val := b[k]
		for i := len(val) - 1; i >= 0; i-- {
			result = append(result, val[i])
		}
	}

	// Append a values.
	for i := len(number) - 1; i >= 0; i-- {
		result = append(result, number[i])
	}

	return result
}
