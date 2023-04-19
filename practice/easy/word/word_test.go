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
	"log"
	"os"
	"testing"
)

// TestFindASubWordGivenCase implements the test given as first example on hackerrank.
func TestFindASubWordGivenCase(t *testing.T) {

	input := []string{"existing pessimist optimist this is"}
	terms := []string{"is"}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "3\n"

	validate(stdout, input, terms)

	result := readTestFile()

	if result != expected {
		t.Errorf("Find a subword first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// readTestFile is responsible for reading the output of the program written in the given writer.
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
