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

// solveFlyRoute calculates the number of valid routes for Jim's flying device.
func solveFlyRoute(arr []int32) int64 {
	routes := int64(0)
	cache := make([]int64, len(arr))

	stack := make([]int32, 0)
	stack = append(stack, 0)

	for i := 1; i < len(arr); i++ {
		// Find next higher sky scrapper to estimate routes.
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
			stack = stack[:len(stack)-1]
		}
		// Count routes until current point.
		if len(stack) > 0 && arr[stack[len(stack)-1]] == arr[i] {
			cache[i] = cache[stack[len(stack)-1]] + 1
			stack = stack[:len(stack)-1]
			routes = routes + cache[i]
		}
		// Push current index.
		stack = append(stack, int32(i))
	}

	return routes * 2
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
//	_, err = strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	arrTemp := strings.Split(readLine(reader), " ")
//
//	var arr []int32
//
//	for arrItr := 0; arrItr < len(arrTemp); arrItr++ {
//		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
//		checkError(err)
//		arrItem := int32(arrItemTemp)
//		arr = append(arr, arrItem)
//	}
//
//	result := solveFlyRoute(arr)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
