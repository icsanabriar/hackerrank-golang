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

// TestCastleGridFirstGivenCase implements the test given as first example on hackerrank.
func TestCastleGridFirstGivenCase(t *testing.T) {
	grid := [][]string{{".", "X", "."}, {".", "X", "."}, {".", ".", "."}}
	startX := int64(0)
	startY := int64(0)
	goalX := int64(0)
	goalY := int64(2)
	visited := make(map[Cell]bool)
	size := int64(3)
	expected := int64(3)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridSecondGivenCase implements the test given as second example on hackerrank.
func TestCastleGridSecondGivenCase(t *testing.T) {
	grid := [][]string{{".", ".", "."}, {".", "X", "."}, {".", "X", "."}}
	startX := int64(2)
	startY := int64(0)
	goalX := int64(0)
	goalY := int64(2)
	visited := make(map[Cell]bool)
	size := int64(3)
	expected := int64(2)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridThirdGivenCase implements the test given as third example on hackerrank.
func TestCastleGridThirdGivenCase(t *testing.T) {
	grid := [][]string{{".", ".", "."}, {".", "X", "."}, {".", "X", "."}}
	startX := int64(2)
	startY := int64(0)
	goalX := int64(2)
	goalY := int64(2)
	visited := make(map[Cell]bool)
	size := int64(3)
	expected := int64(3)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridFourthCase implements the test given as fourth example on hackerrank.
func TestCastleGridFourthCase(t *testing.T) {
	grid := [][]string{{"X", ".", "X"}, {".", ".", "."}, {".", "X", "."}}
	startX := int64(2)
	startY := int64(2)
	goalX := int64(1)
	goalY := int64(0)
	visited := make(map[Cell]bool)
	size := int64(3)
	expected := int64(2)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid fourth case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCastleGridEdgeCase implements the test given as edge example on hackerrank.
func TestCastleGridEdgeCase(t *testing.T) {
	grid := [][]string{{".", ".", "."}, {"X", "X", "X"}, {".", "X", "."}}
	startX := int64(2)
	startY := int64(0)
	goalX := int64(2)
	goalY := int64(2)
	visited := make(map[Cell]bool)
	size := int64(3)
	expected := int64(-1)

	result := minimumMoves(grid, startX, startY, goalX, goalY, visited, size)

	if result != expected {
		t.Errorf("Castle grid edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
