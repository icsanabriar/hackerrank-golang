// Copyright 2023 Iván Camilo Sanabria
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
	"testing"
)

// TestDiagonalDifferenceFirstGivenCase implements the test given as first example on hackerrank.
func TestDiagonalDifferenceFirstGivenCase(t *testing.T) {

	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	expected := int32(2)

	result := diagonalDifference(arr)

	if result != expected {
		t.Errorf("Diagonal difference first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestDiagonalDifferenceSecondGivenCase implements the test given as second example on hackerrank.
func TestDiagonalDifferenceSecondGivenCase(t *testing.T) {

	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	expected := int32(15)

	result := diagonalDifference(arr)

	if result != expected {
		t.Errorf("Diagonal difference second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
