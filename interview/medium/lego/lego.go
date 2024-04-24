// Copyright 2023 Iván Camilo Sanabria
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

// M used to reduce the result.
const M = 1000000007

// legoBlocks estimates the number of valid wall formations modulo M.
func legoBlocks(h int, w int) int64 {
	splits := countSplits()

	caches := make([]int64, w+1)
	for i := 0; i <= w; i++ {
		caches[i] = powMod(splits[i], int64(h))
	}

	sConstr := make([]int64, w+1)
	sConstr[1] = 1

	for i := 2; i <= w; i++ {
		sConstr[i] = caches[i]
		for j := 1; j < i; j++ {
			sConstr[i] = subtractMod(sConstr[i], multiplyMod(sConstr[j], caches[i-j]))
		}
	}

	return sConstr[w]
}

// addMod adds a and b and then apply module M.
func addMod(a, b int64) int64 {
	return (a + b) % M
}

// multiplyMod multiply a and b and then apply module M.
func multiplyMod(a, b int64) int64 {
	return a * b % M
}

// subtractMod subtract a and b and then apply module M.
func subtractMod(a, b int64) int64 {
	return (a + M - b) % M
}

// powMod pow a and b and then apply module M.
func powMod(a, b int64) int64 {
	res := int64(1)
	base := a % M

	for b != 0 {
		if b%2 == 1 {
			res = (res * base) % M
		}
		base = (base * base) % M
		b >>= 1
	}

	return res % M
}

// countSplits estimates the number of splits on 1001 rows.
func countSplits() []int64 {
	splits := make([]int64, 1001)
	for i := 1; i <= 4; i++ {
		splits[i] = 1
	}

	for i := 1; i < 1001; i++ {
		for j := 1; j <= 4; j++ {
			if i-j >= 0 {
				splits[i] = addMod(splits[i], splits[i-j])
			}
		}
	}

	return splits
}
