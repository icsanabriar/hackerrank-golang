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

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			fmt.Println("Error reading output path file.")
		}
	}(stdout)

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		a, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)

		b, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		c, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)

		d, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)

		e, err := strconv.ParseInt(firstMultipleInput[4], 10, 64)
		checkError(err)

		f, err := strconv.ParseInt(firstMultipleInput[5], 10, 64)
		checkError(err)

		g, err := strconv.ParseInt(firstMultipleInput[6], 10, 64)
		checkError(err)

		h, err := strconv.ParseInt(firstMultipleInput[7], 10, 64)
		checkError(err)

		n, err := strconv.ParseInt(firstMultipleInput[8], 10, 64)
		checkError(err)

		result := recurrence(a, b, c, d, e, f, g, h, n)
		printResult(result, writer)
	}

	_ = writer.Flush()
}

// printResult function provided to simplify the output process.
func printResult(result []int64, writer *bufio.Writer) {
	for i, resultItem := range result {
		_, _ = fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			_, _ = fmt.Fprintf(writer, " ")
		}
	}

	_, _ = fmt.Fprintf(writer, "\n")
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
