// Copyright 2020 Iván Camilo Sanabria
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

	nm := strings.Split(readLine(reader), " ")

	n, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)

	m, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)

	var operations [][]int64
	for operationsRowItr := 0; operationsRowItr < int(m); operationsRowItr++ {
		operationsRowTemp := strings.Split(readLine(reader), " ")

		var operationsRow []int64
		for _, operationsRowItem := range operationsRowTemp {
			operationsItem, err := strconv.ParseInt(operationsRowItem, 10, 64)
			checkError(err)
			operationsRow = append(operationsRow, operationsItem)
		}

		if len(operationsRow) != 3 {
			panic("Bad input")
		}

		operations = append(operations, operationsRow)
	}

	result := estimate(n, operations)

	_, _ = fmt.Fprintf(writer, "%d\n", result)

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
