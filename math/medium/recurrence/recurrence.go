// Copyright 2021 Iván Camilo Sanabria
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

// MODULUS is used to return output values.
const MODULUS = 1000000000

// identical build the identity matrix for the given size.
func identical(size int) [][]int64 {
	result := make([][]int64, size)

	for i := range result {
		result[i] = make([]int64, size)
	}

	for i := 0; i < size; i++ {
		result[i][i] = 1
	}

	return result
}

// add values of integers and normalized with the given MODULUS.
func add(x int64, y int64) int64 {
	return (x + y) % MODULUS
}

// multiply values of integers and normalized with the given MODULUS.
func multiply(x int64, y int64) int64 {
	return x * y % MODULUS
}

// pow numbers of integers on the given matrix.
func powMatrix(base [][]int64, n int64) [][]int64 {
	size := len(base[0])
	result := identical(size)

	for n != 0 {
		if (n & 1) != 0 {
			result = multiplyMatrix(result, base)
		}

		base = multiplyMatrix(base, base)
		n >>= 1
	}

	return result
}

// multiplyVectors multiply vector with matrix based on initial and transition states.
func multiplyVectors(initial []int64, transition [][]int64) []int64 {
	result := make([]int64, len(initial))

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			result[i] = add(result[i], multiply(initial[j], transition[j][i]))
		}
	}

	return result
}

// multiplyMatrix multiply matrix with matrix based on the given values.
func multiplyMatrix(a [][]int64, b [][]int64) [][]int64 {
	size := len(a[0])
	result := make([][]int64, size)

	for i := range result {
		result[i] = make([]int64, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] = add(result[i][j], multiply(a[i][k], b[k][j]))
			}
		}
	}

	return result
}

// transitionState build transition state matrix of X and Y using start values.
func transitionState(a int32, b int32, c int32, d int32, e int32, f int32, g int32, h int32) [][]int64 {
	matrix := make([][]int64, 22)

	for i := range matrix {
		matrix[i] = make([]int64, 22)
	}

	matrix[a-1][0] = 1
	matrix[b+10][0]++
	matrix[c+10][0]++
	matrix[9][0] = 1

	for i := 1; i < 9; i++ {
		matrix[i-1][i] = 1
	}

	matrix[9][9] = int64(d)
	matrix[10][9] = int64(d)

	matrix[10][10] = int64(d)

	matrix[e+10][11] = 1
	matrix[f-1][11]++
	matrix[g-1][11]++
	matrix[20][11] = 1

	for i := 12; i < 20; i++ {
		matrix[i-1][i] = 1
	}

	matrix[20][20] = int64(h)
	matrix[21][20] = int64(h)

	matrix[21][21] = int64(h)

	return matrix
}

// bootstrap state in order to begin calculation.
func bootstrap() []int64 {
	state := make([]int64, 22)

	for i := 0; i < 9; i++ {
		state[i] = 1
	}

	for i := 10; i < 20; i++ {
		state[i] = 1
	}

	state[21] = 1

	return state
}

// recurrence calculates the value of Xi and Yi based on the given parameters and value of n.
func recurrence(a int32, b int32, c int32, d int32, e int32, f int32, g int32, h int32, n int64) []int64 {
	initial := bootstrap()
	transition := transitionState(a, b, c, d, e, f, g, h)
	state := multiplyVectors(initial, powMatrix(transition, n+1))
	return []int64{state[0], state[11]}
}
