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

import "testing"

// TestClosestNumberFirstGivenCase implements the test given as first example on hackerrank.
func TestClosestNumberFirstGivenCase(t *testing.T) {
	a := int64(349)
	b := int64(1)
	x := int64(4)
	expected := int64(348)

	result := closestNumber(a, b, x)

	if result != expected {
		t.Errorf("Closest Number first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestClosestNumberSecondGivenCase implements the test given as second example on hackerrank.
func TestClosestNumberSecondGivenCase(t *testing.T) {
	a := int64(395)
	b := int64(1)
	x := int64(7)
	expected := int64(392)

	result := closestNumber(a, b, x)

	if result != expected {
		t.Errorf("Closest Number second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestClosestNumberThirdGivenCase implements the test given as third example on hackerrank.
func TestClosestNumberThirdGivenCase(t *testing.T) {
	a := int64(4)
	b := int64(-2)
	x := int64(2)
	expected := int64(0)

	result := closestNumber(a, b, x)

	if result != expected {
		t.Errorf("Closest Number third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestClosestNumberEdgeCase implements the test given as edge case to increase coverage.
func TestClosestNumberEdgeCase(t *testing.T) {
	a := int64(1)
	b := int64(0)
	x := int64(2)
	expected := int64(0)

	result := closestNumber(a, b, x)

	if result != expected {
		t.Errorf("Closest Number edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestClosestNumberHiddenCase implements the test given as hidden case to increase coverage.
func TestClosestNumberHiddenCase(t *testing.T) {
	a := int64(1)
	b := int64(973594325)
	x := int64(1)
	expected := int64(1)

	result := closestNumber(a, b, x)

	if result != expected {
		t.Errorf("Closest Number hidden case was incorrect, got: %d, want: %d.", result, expected)
	}
}
