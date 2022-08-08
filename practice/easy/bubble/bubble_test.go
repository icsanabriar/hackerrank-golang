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

import "testing"

// TestBubbleSortFirstGivenCase implements the test given as first example on hackerrank.
func TestBubbleSortFirstGivenCase(t *testing.T) {
	input := []int32{1, 2, 3}
	expected := "Array is sorted in 0 swaps.\nFirst Element: 1\nLast Element: 3\n"

	result := countSwaps(input)

	if result != expected {
		t.Errorf("Bubble sort first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestBubbleSortSecondGivenCase implements the test given as first example on hackerrank.
func TestBubbleSortSecondGivenCase(t *testing.T) {
	input := []int32{3, 2, 1}
	expected := "Array is sorted in 3 swaps.\nFirst Element: 1\nLast Element: 3\n"

	result := countSwaps(input)

	if result != expected {
		t.Errorf("Bubble sort second case was incorrect, got: %s, want: %s.", result, expected)
	}
}
