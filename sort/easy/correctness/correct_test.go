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
	"log"
	"os"
	"testing"
)

// TestCorrectnessLoopInvariantGivenCase implements the test given as first example on hackerrank.
func TestCorrectnessLoopInvariantGivenCase(t *testing.T) {

	arr := []int32{7, 4, 3, 5, 6, 2}
	expected := "2 3 4 5 6 7 \n"

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	insertionSort(stdout, arr)

	result := readTestFile()

	if result != expected {
		t.Errorf("Correctness loop invariant first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// readTestFile is responsible of reading the output of the program written in the given writer.
func readTestFile() string {

	text := ""

	file, err := os.Open(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			fmt.Println("Error reading output path file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
