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

import "testing"

// TestGlassesFirstGivenCase implements the test given as first example on hackerrank.
func TestGlassesFirstGivenCase(t *testing.T) {

	input := [][]int32{{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}

	expected := int32(19)

	result := hourglassSum(input)

	if result != expected {
		t.Errorf("2D Array - DS first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestGlassesSecondGivenCase implements the test given as second example on hackerrank.
func TestGlassesSecondGivenCase(t *testing.T) {

	input := [][]int32{{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0}}

	expected := int32(28)

	result := hourglassSum(input)

	if result != expected {
		t.Errorf("2D Array - DS second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
