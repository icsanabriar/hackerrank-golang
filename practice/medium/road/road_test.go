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

// TestRoadsAndLibrariesFirstGivenCase implements the test given as first example on hackerrank.
func TestRoadsAndLibrariesFirstGivenCase(t *testing.T) {

	n := int32(3)
	lib := int32(2)
	road := int32(1)

	cities := [][]int32{{1, 2}, {3, 1}, {2, 3}}
	expected := int64(4)

	result := roadsAndLibraries(n, lib, road, cities)

	if result != expected {
		t.Errorf("Road and libraries first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestRoadsAndLibrariesSecondGivenCase implements the test given as second example on hackerrank.
func TestRoadsAndLibrariesSecondGivenCase(t *testing.T) {

	n := int32(6)
	lib := int32(2)
	road := int32(5)

	cities := [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}
	expected := int64(12)

	result := roadsAndLibraries(n, lib, road, cities)

	if result != expected {
		t.Errorf("Road and libraries second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
