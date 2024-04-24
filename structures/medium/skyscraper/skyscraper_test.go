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

// TestSkyscraperFirstGivenCase implements the test given as first example on hackerrank.
func TestSkyscraperFirstGivenCase(t *testing.T) {
	input := []int64{3, 2, 1, 2, 3, 3}
	expected := int64(8)

	result := solveFlyRoute(input)

	if result != expected {
		t.Errorf("Jim and the skyscrapers first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestSkyscraperSecondGivenCase implements the test given as second example on hackerrank.
func TestSkyscraperSecondGivenCase(t *testing.T) {
	input := []int64{1, 1000, 1}
	expected := int64(0)

	result := solveFlyRoute(input)

	if result != expected {
		t.Errorf("Jim and the skyscrapers second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestSkyscraperEdgeCase implements the test given as edge example.
func TestSkyscraperEdgeCase(t *testing.T) {
	input := []int64{2, 2, 2, 2, 2, 2, 1, 1, 2, 2, 1, 2, 1, 1, 1, 1, 1, 1, 2, 2}
	expected := int64(142)

	result := solveFlyRoute(input)

	if result != expected {
		t.Errorf("Jim and the skyscrapers edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
