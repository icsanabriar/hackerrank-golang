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

	n, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)

	m, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	k, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)

	var track [][]int64
	for i := 0; i < int(k); i++ {
		trackRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var trackRow []int64
		for _, trackRowItem := range trackRowTemp {
			trackItem, err := strconv.ParseInt(trackRowItem, 10, 64)
			checkError(err)
			trackRow = append(trackRow, trackItem)
		}

		if len(trackRow) != 3 {
			panic("Bad input")
		}

		track = append(track, trackRow)
	}

	result := gridlandMetro(n, m, track)

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
