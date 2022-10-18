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

// TestSplitPhoneNumberGivenCase implements the test given as first example on hackerrank.
func TestSplitPhoneNumberGivenCase(t *testing.T) {

	input := []string{"1 877 2638277", "91-011-23413627"}
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "CountryCode=1,LocalAreaCode=877,Number=2638277\n" +
		"CountryCode=91,LocalAreaCode=011,Number=23413627\n"

	classify(stdout, input)

	result := readTestFile()

	if result != expected {
		t.Errorf("Split phone number first case was incorrect, got: %s, want: %s.", result, expected)
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
