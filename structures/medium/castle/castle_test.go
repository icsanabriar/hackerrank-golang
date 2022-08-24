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

// TestCastleGridFirstGivenCase implements the test given as first example on hackerrank.
func TestCastleGridFirstGivenCase(t *testing.T) {

	grid := [][]string{{".", "X", "."}, {".", "X", "."}, {".", ".", "."}}
	startX := int32(0)
	startY := int32(0)
	goalX := int32(0)
	goalY := int32(2)
	visited := make(map[Cell]bool, 0)
	size := int32(3)
	expected := int32(3)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridSecondGivenCase implements the test given as second example on hackerrank.
func TestCastleGridSecondGivenCase(t *testing.T) {

	grid := [][]string{{".", ".", "."}, {".", "X", "."}, {".", "X", "."}}
	startX := int32(2)
	startY := int32(0)
	goalX := int32(0)
	goalY := int32(2)
	visited := make(map[Cell]bool, 0)
	size := int32(3)
	expected := int32(2)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridThirdGivenCase implements the test given as third example on hackerrank.
func TestCastleGridThirdGivenCase(t *testing.T) {

	grid := [][]string{{".", ".", "."}, {".", "X", "."}, {".", "X", "."}}
	startX := int32(2)
	startY := int32(0)
	goalX := int32(2)
	goalY := int32(2)
	visited := make(map[Cell]bool, 0)
	size := int32(3)
	expected := int32(3)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
