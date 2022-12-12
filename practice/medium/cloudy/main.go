// Copyright 2022 Iván Camilo Sanabria
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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int64

	for i := 0; i < int(n); i++ {
		pItem, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		p = append(p, pItem)
	}

	xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var x []int64

	for i := 0; i < int(n); i++ {
		xItem, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		x = append(x, xItem)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	yTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var y []int64

	for i := 0; i < int(m); i++ {
		yItem, err := strconv.ParseInt(yTemp[i], 10, 64)
		checkError(err)
		y = append(y, yItem)
	}

	rTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var r []int64

	for i := 0; i < int(m); i++ {
		rItem, err := strconv.ParseInt(rTemp[i], 10, 64)
		checkError(err)
		r = append(r, rItem)
	}

	result := maximumPeople(p, x, y, r)

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
