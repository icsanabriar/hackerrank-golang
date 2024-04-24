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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n1, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)

	n2, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	n3, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)

	h1Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h1 []int64
	for i := 0; i < int(n1); i++ {
		h1Item, err := strconv.ParseInt(h1Temp[i], 10, 64)
		checkError(err)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h2 []int64
	for i := 0; i < int(n2); i++ {
		h2Item, err := strconv.ParseInt(h2Temp[i], 10, 64)
		checkError(err)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h3 []int64
	for i := 0; i < int(n3); i++ {
		h3Item, err := strconv.ParseInt(h3Temp[i], 10, 64)
		checkError(err)
		h3 = append(h3, h3Item)
	}

	result := equalStacks(h1, h2, h3)

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
