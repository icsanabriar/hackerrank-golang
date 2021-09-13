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

//dynamicArray uses the given queries to complete arrays using the given n.
func dynamicArray(n int32, queries [][]int32) []int32 {
	last := int32(0)
	results := make([]int32, 0)
	dynamic := make([][]int32, n)

	for _, item := range queries {
		operation := item[0]
		x := item[1]
		y := item[2]

		index := (x ^ last) % n

		if 1 == operation {
			array := dynamic[index]
			dynamic[index] = append(array, y)
		} else if 2 == operation {
			array := dynamic[index]
			index = y % int32(len(array))
			last = array[index]
			results = append(results, last)
		}
	}

	return results
}

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
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
//	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	q := int32(qTemp)
//
//	var queries [][]int32
//	for i := 0; i < int(q); i++ {
//		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")
//
//		var queriesRow []int32
//		for _, queriesRowItem := range queriesRowTemp {
//			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
//			checkError(err)
//			queriesItem := int32(queriesItemTemp)
//			queriesRow = append(queriesRow, queriesItem)
//		}
//
//		if len(queriesRow) != 3 {
//			panic("Bad input")
//		}
//
//		queries = append(queries, queriesRow)
//	}
//
//	result := dynamicArray(n, queries)
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
