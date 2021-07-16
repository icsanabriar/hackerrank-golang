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

import "sort"

// cutTheSticks return the number of sticks after each iteration of cutting.
func cutTheSticks(arr []int32) []int32 {

	cache := make(map[int32]int32)

	for i := range arr {
		length := arr[i]
		cache[length]++
	}

	keys := make([]int32, 0, len(cache))

	for k := range cache {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	var result []int32
	left := int32(len(arr))

	for _, k := range keys {
		result = append(result, left)
		left = left - cache[k]
	}

	return result
}

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
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
//	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
//
//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var arr []int32
//
//	for i := 0; i < int(n); i++ {
//		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//		checkError(err)
//		arrItem := int32(arrItemTemp)
//		arr = append(arr, arrItem)
//	}
//
//	result := cutTheSticks(arr)
//
//	for i, resultItem := range result {
//		_, _ = fmt.Fprintf(writer, "%d", resultItem)
//
//		if i != len(result) - 1 {
//			_, _ = fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
