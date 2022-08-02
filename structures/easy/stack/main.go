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
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			fmt.Println("Error reading output path file.")
		}
	}(stdout)

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var s = stack{}
	var maxStack = stack{}

	for i := 0; i < int(n); i++ {
		command := strings.Split(readLine(reader), " ")

		// Command 1 is the only two parameters.
		if len(command) == 2 {

			eTemp, err := strconv.ParseInt(command[1], 10, 64)
			checkError(err)
			e := int32(eTemp)

			// Push the element to the stacks.
			s.push(e)
			maxStack.pushMax(e)

		} else {

			// Check command 2 to delete top element from the stack and cache.
			if command[0] == "2" {
				s.pop()
				maxStack.pop()
			} else {
				// Print the maximum value using the cache.
				_, _ = fmt.Fprintf(writer, "%d\n", maxStack.peek())
			}
		}
	}

	_ = writer.Flush()
}

// readLine function provided by hacker rank to execute the code on platform.
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// checkError function provided by hacker rank to execute the code on platform.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
