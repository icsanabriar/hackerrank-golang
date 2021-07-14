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
	"sort"
)

// repeatedString return an integer representing the number of occurrences of 'a' in the length n.
func repeatedString(s string, n int64) int64 {

	searchChar := string('a')
	cacheChar := make([]int, 0)
	repeated := int64(0)

	for i, char := range s {
		if string(char) == searchChar {
			cacheChar = append(cacheChar, i)
		}
	}

	// Full repetitions  of the string.
	repetitions := n / int64(len(s))
	repeated = repetitions * int64(len(cacheChar))

	// Last part if there is a split of the string.
	split := n % int64(len(s))

	if split != 0 {
		sort.Ints(cacheChar)
		broken := false

		for i := range cacheChar {
			if split <= int64(cacheChar[i]) {
				repeated = repeated + int64(i)
				broken = true
				break
			}
		}

		if !broken {
			repeated = repeated + int64(len(cacheChar))
		}
	}

	return repeated
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
//	s := readLine(reader)
//
//	n, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	result := repeatedString(s, n)
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
