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

// maxN is the maximum value of N.
const maxN = int32(1000000)

// cache is a map to cache steps to 0.
var cache = buildCache()

// min returns the minor of a or b.
func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

// buildCache generates an array to memorize the steps required to down to 0 specific index of the array.
func buildCache() []int32 {

	temp := make([]int32, maxN+1)

	for i := int32(1); i < maxN; i++ {
		temp[i] = temp[i-1] + 1

		for j := int32(2); j*j <= i; j++ {

			if i%j == 0 {
				temp[i] = min(temp[i], temp[i/j]+1)
			}
		}
	}

	return temp
}

// downToZero reads the given n value from cache and return it.
func downToZero(n int32) int32 {
	return cache[n]
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
//	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	q := int32(qTemp)
//
//	for qItr := 0; qItr < int(q); qItr++ {
//
//		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		n := int32(nTemp)
//
//		result := downToZero(n)
//
//		_, _ = fmt.Fprintf(writer, "%d\n", result)
//	}
//
//	_ = writer.Flush()
//}
