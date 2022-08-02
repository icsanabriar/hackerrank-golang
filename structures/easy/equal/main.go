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

	n1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n1 := int32(n1Temp)

	n2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n2 := int32(n2Temp)

	n3Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n3 := int32(n3Temp)

	h1Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h1 []int32

	for i := 0; i < int(n1); i++ {
		h1ItemTemp, err := strconv.ParseInt(h1Temp[i], 10, 64)
		checkError(err)
		h1Item := int32(h1ItemTemp)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h2 []int32

	for i := 0; i < int(n2); i++ {
		h2ItemTemp, err := strconv.ParseInt(h2Temp[i], 10, 64)
		checkError(err)
		h2Item := int32(h2ItemTemp)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h3 []int32

	for i := 0; i < int(n3); i++ {
		h3ItemTemp, err := strconv.ParseInt(h3Temp[i], 10, 64)
		checkError(err)
		h3Item := int32(h3ItemTemp)
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
