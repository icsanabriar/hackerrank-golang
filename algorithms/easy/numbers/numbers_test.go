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

import "testing"

// TestNumbersFirstGivenCase implements the test given as first example on hackerrank.
func TestNumbersFirstGivenCase(t *testing.T) {
	input := []int64{4, 6, 5, 3, 3, 1}
	expected := int64(3)

	result := pickingNumbers(input)

	if result != expected {
		t.Errorf("Picking numbers first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestNumbersSecondGivenCase implements the test given as second example on hackerrank.
func TestNumbersSecondGivenCase(t *testing.T) {
	input := []int64{1, 2, 2, 3, 1, 2}
	expected := int64(5)

	result := pickingNumbers(input)

	if result != expected {
		t.Errorf("Picking numbers second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestNumbersBasicGivenCase implements the test given as basic example on hackerrank.
func TestNumbersBasicGivenCase(t *testing.T) {
	input := []int64{100}
	expected := int64(1)

	result := pickingNumbers(input)

	if result != expected {
		t.Errorf("Picking numbers basic case was incorrect, got: %d, want: %d.", result, expected)
	}
}
