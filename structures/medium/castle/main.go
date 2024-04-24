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

	n, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)

	var grid []string
	for i := 0; i < n; i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	startX, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)

	startY, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	goalX, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)

	goalY, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)

	parsed := make([][]string, 0, len(grid))
	for i := range grid {
		row := strings.Split(grid[i], "")
		parsed = append(parsed, row)
	}

	visited := make(map[Cell]bool)
	result := minimumMoves(parsed, startX, startY, goalX, goalY, visited, int64(n))

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
