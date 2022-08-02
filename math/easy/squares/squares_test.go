// Copyright 2021 Iván Camilo Sanabria
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

// TestCuttingPaperSquaresGivenCase implements the test given as first example on hackerrank.
func TestCuttingPaperSquaresGivenCase(t *testing.T) {

	n := int64(3)
	m := int64(1)
	expected := int64(2)

	result := cutPaper(n, m)

	if result != expected {
		t.Errorf("Cutting Paper Squares first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCuttingPaperSquaresEdgeCase implements the edge case to validate results on hackerrank.
func TestCuttingPaperSquaresEdgeCase(t *testing.T) {

	n := int64(3)
	m := int64(2)
	expected := int64(5)

	result := cutPaper(n, m)

	if result != expected {
		t.Errorf("Cutting Paper Squares edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCuttingPaperSquaresHiddenCase implements the hidden case to validate results on hackerrank.
func TestCuttingPaperSquaresHiddenCase(t *testing.T) {

	n := int64(689715240)
	m := int64(759842301)
	expected := int64(524074814996367239)

	result := cutPaper(n, m)

	if result != expected {
		t.Errorf("Cutting Paper Squares edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
