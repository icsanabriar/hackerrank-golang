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

// TestRectangleFirstGivenCase implements the test given as first example on hackerrank.
func TestRectangleFirstGivenCase(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	expected := int64(9)

	result := largestRectangle(input)

	if result != expected {
		t.Errorf("Largest rectangle first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRectangleSecondGivenCase implements the test given as second example on hackerrank.
func TestRectangleSecondGivenCase(t *testing.T) {
	input := []int64{1, 3, 5, 9, 11}
	expected := int64(18)

	result := largestRectangle(input)

	if result != expected {
		t.Errorf("Largest rectangle second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRectangleThirdGivenCase implements the test given as third example on hackerrank.
func TestRectangleThirdGivenCase(t *testing.T) {
	input := []int64{11, 11, 10, 10, 10}
	expected := int64(50)

	result := largestRectangle(input)

	if result != expected {
		t.Errorf("Largest rectangle third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
