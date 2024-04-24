// Copyright 2022 Iván Camilo Sanabria
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
	"io"
	"os"
	"strconv"
	"strings"
)

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	input := readInput(reader, n)
	countSort(os.Stdout, input)
}

// readInput reads the given input to process information in order.
func readInput(reader *bufio.Reader, n int64) map[int64]string {
	helper := make(map[int64]string)

	for i := int64(0); i < (n / 2); i++ {
		row := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		if len(row) != 2 {
			panic("Bad input")
		}

		key, _ := strconv.ParseInt(row[0], 10, 64)
		value := "-"

		if values, ok := helper[key]; ok {
			helper[key] = values + " " + value
		} else {
			helper[key] = value
		}
	}

	for i := n / 2; i < n; i++ {
		row := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		if len(row) != 2 {
			panic("Bad input")
		}

		key, _ := strconv.ParseInt(row[0], 10, 64)
		value := row[1]

		if values, ok := helper[key]; ok {
			helper[key] = values + " " + value
		} else {
			helper[key] = value
		}
	}

	return helper
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
