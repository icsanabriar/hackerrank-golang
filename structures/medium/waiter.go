// Copyright 2020 Iván Camilo Sanabria
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

import (
	"math/big"
	"sort"
)

// prime calculates the i-th prime using a starting int.
func prime(previous int32) int32 {
	for i := previous + 1; ; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			return i
		}
	}
}

// waiter separates the given array of plates (number) into Ai Bi piles doing q iterations.
// Returns an array of int with the order of the piles B0, B1, .. Bi, Ai taking into account
// the TOP element which is the last one of the pile.
func waiter(number []int32, q int) []int32 {

	a := make(map[int][]int32)
	b := make(map[int][]int32)
	p := int32(2)

	// Build piles from ai.
	for i := 1; i <= q; i++ {
		for j := len(number) - 1; j >= 0; j-- {
			e := number[j]
			if e%p == 0 {
				if val, ok := b[i]; ok {
					b[i] = append(val, e)
				} else {
					b[i] = []int32{e}
				}
			} else {
				if val, ok := a[i]; ok {
					a[i] = append(val, e)
				} else {
					a[i] = []int32{e}
				}
			}
		}
		number = a[i]
		delete(a, i)
		p = prime(p)
	}

	result := make([]int32, 0)

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

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer func(stdout *os.File) {
//		err := stdout.Close()
//		if err != nil {
//			fmt.Println("Error reading output path file.")
//		}
//	}(stdout)
//
//	writer := bufio.NewWriterSize(stdout, 1024*1024)
//
//	nq := strings.Split(readLine(reader), " ")
//
//	nTemp, err := strconv.ParseInt(nq[0], 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	qTemp, err := strconv.ParseInt(nq[1], 10, 64)
//	checkError(err)
//	q := int(qTemp)
//
//	numberTemp := strings.Split(readLine(reader), " ")
//
//	var number []int32
//
//	for numberItr := 0; numberItr < int(n); numberItr++ {
//		numberItemTemp, err := strconv.ParseInt(numberTemp[numberItr], 10, 64)
//		checkError(err)
//		numberItem := int32(numberItemTemp)
//		number = append(number, numberItem)
//	}
//
//	result := waiter(number, q)
//
//	for resultItr, resultItem := range result {
//		_, _ = fmt.Fprintf(writer, "%d", resultItem)
//
//		if resultItr != len(result)-1 {
//			_, _ = fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
