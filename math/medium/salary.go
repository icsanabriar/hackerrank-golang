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
	"math"
)

// gcd calculate the great common divisor of given a,b.
func gcd(a int64, b int64) int64 {

	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// diffGcd calculate gcd of the absolute difference of 2 elements in position i and i-1.
func diffGcd(a []int64) int64 {
	result := int64(math.Abs(float64(a[1] - a[0])))

	for i := 2; i < len(a); i++ {
		result = gcd(result, int64(math.Abs(float64(a[i]-a[i-1]))))
	}

	return result
}

// normalize finds the common salaries for the given array a, taking into account the given queries
// that represents raises for all employees of a before normalization is executed.
func normalize(a []int64, queries []int32) []int64 {

	result := make([]int64, 0)
	diff := diffGcd(a)

	for i := 0; i < len(queries); i++ {
		k := int64(queries[i])
		result = append(result, gcd(diff, a[0]+k))
	}

	return result
}

// main function provided by hacker rank to execute the code on platform.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)
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
//	q := int32(qTemp)
//
//	aTemp := strings.Split(readLine(reader), " ")
//
//	var a []int64
//
//	fmt.Printf("Values for n %d and len(aTemp) %d.", n, len(aTemp))
//
//	for i := 0; i < len(aTemp); i++ {
//		aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
//		checkError(err)
//		a = append(a, aItem)
//	}
//
//	var queries []int32
//
//	for i := 0; i < int(q); i++ {
//		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		queriesItem := int32(queriesItemTemp)
//		queries = append(queries, queriesItem)
//	}
//
//	result := normalize(a, queries)
//
//	for i, resultItem := range result {
//		_, _ = fmt.Fprintf(writer, "%d", resultItem)
//
//		if i != len(result)-1 {
//			_, _ = fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
