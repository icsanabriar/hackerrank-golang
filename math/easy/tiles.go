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

import (
	"math"
)

// movingTiles calculates the area of intersection square between two tiles.
func movingTiles(l int32, s1 int32, s2 int32, queries []int64) []float64 {

	results := make([]float64, 0)

	for i := range queries {
		q := (float64(l)*math.Sqrt2 - math.Sqrt(float64(queries[i]))*math.Sqrt2) / math.Abs(float64(s1)-float64(s2))
		results = append(results, q)
	}

	return results
}

// main function provided by hacker rank to execute the code on platform.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
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
//	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
//
//	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	lTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//	checkError(err)
//	l := int32(lTemp)
//
//	s1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	s1 := int32(s1Temp)
//
//	s2Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
//	checkError(err)
//	s2 := int32(s2Temp)
//
//	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	var queries []int64
//
//	for i := 0; i < int(queriesCount); i++ {
//		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//		checkError(err)
//		queries = append(queries, queriesItemTemp)
//	}
//
//	result := movingTiles(l, s1, s2, queries)
//
//	for i, resultItem := range result {
//		_, _ = fmt.Fprintf(writer, "%f", resultItem)
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
