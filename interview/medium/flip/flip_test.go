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

import "testing"

// TestFlippingMatrixFirstGivenCase implements the test given as first example on hackerrank.
func TestFlippingMatrixFirstGivenCase(t *testing.T) {

	matrix := [][]int32{{112, 42, 83, 119}, {56, 125, 56, 49}, {15, 78, 101, 43}, {62, 98, 114, 108}}
	expected := int32(414)

	result := flippingMatrix(matrix)

	if result != expected {
		t.Errorf("Flipping matrix first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestFlippingMatrixSecondGivenCase implements the test given as second example on hackerrank.
func TestFlippingMatrixSecondGivenCase(t *testing.T) {

	matrix := [][]int32{{107, 54, 128, 15}, {12, 75, 110, 138}, {100, 96, 34, 85}, {75, 15, 28, 112}}
	expected := int32(488)

	result := flippingMatrix(matrix)

	if result != expected {
		t.Errorf("Flipping matrix second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
