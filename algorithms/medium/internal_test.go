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
	"errors"
	"fmt"
	"strings"
	"testing"
)

// TestReadEmptyLine implements the test case to read empty line using the given buffer reader.
func TestReadEmptyLine(t *testing.T) {

	in := strings.NewReader("")
	reader := bufio.NewReader(in)

	result := readLine(reader)
	expected := ""

	if result != expected {
		t.Errorf("Read line case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestReadStringLine implements the test case to read not empty line using the given buffer reader.
func TestReadStringLine(t *testing.T) {

	in := strings.NewReader("{...sample of line...}")
	reader := bufio.NewReader(in)

	result := readLine(reader)
	expected := "{...sample of line...}"

	if result != expected {
		t.Errorf("Read line case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestCheckNilError implements the test case to check error with nil value and the program should not panic.
func TestCheckNilError(t *testing.T) {
	checkError(nil)
}

// TestCheckError implements the test case to check given error and the program should panic.
func TestCheckError(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in check error function", r)
		}
	}()

	input := errors.New("error to test check error function")
	checkError(input)

	t.Errorf("The check error function did not panic.")
}
