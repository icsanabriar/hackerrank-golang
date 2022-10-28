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

// TestGridlandMetroGivenCase implements the test given as first example on hackerrank.
func TestGridlandMetroGivenCase(t *testing.T) {

	n := int32(4)
	m := int32(4)

	track := [][]int32{{2, 2, 3}, {3, 1, 4}, {4, 4, 4}}
	expected := int64(9)

	result := gridlandMetro(n, m, track)

	if result != expected {
		t.Errorf("Gridland metro first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestGridlandMetroHiddenCase implements the test given as hidden example on hackerrank.
func TestGridlandMetroHiddenCase(t *testing.T) {

	n := int32(1)
	m := int32(5)

	track := [][]int32{{1, 1, 2}, {1, 2, 4}, {1, 3, 5}}
	expected := int64(0)

	result := gridlandMetro(n, m, track)

	if result != expected {
		t.Errorf("Gridland metro hidden case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestGridlandMetroEdgeCase implements the test given as edge case.
func TestGridlandMetroEdgeCase(t *testing.T) {

	n := int32(1)
	m := int32(10)

	track := [][]int32{{1, 1, 2}, {1, 3, 4}, {1, 8, 9}}
	expected := int64(4)

	result := gridlandMetro(n, m, track)

	if result != expected {
		t.Errorf("Gridland metro edge case was incorrect, got: %v, want: %v.", result, expected)
	}
}
