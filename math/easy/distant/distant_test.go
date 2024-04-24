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

import (
	"math/big"
	"testing"
)

// TestMostDistantFirstGivenCase implements the test given as first example on hackerrank.
func TestMostDistantFirstGivenCase(t *testing.T) {
	input := [][]int64{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	expected := float64(2)

	solution := solve(input)

	result := big.NewFloat(solution).Cmp(big.NewFloat(expected))

	if 0 != result {
		t.Errorf("Most distant first case was incorrect, got: %f, want: %f.", solution, expected)
	}
}

// TestMostDistantSecondGivenCase implements the test given as second example on hackerrank.
func TestMostDistantSecondGivenCase(t *testing.T) {
	input := [][]int64{{0, -5}, {-7, 0}, {0, -6}, {-4, 0}, {0, 0}}
	expected := 9.219544457292887

	solution := solve(input)

	result := big.NewFloat(solution).Cmp(big.NewFloat(expected))

	if 0 != result {
		t.Errorf("Most distant second case was incorrect, got: %f, want: %f.", solution, expected)
	}
}

// TestMostDistantEdgeXCase implements the test given as edge x example on hackerrank.
func TestMostDistantEdgeXCase(t *testing.T) {
	input := [][]int64{{0, 0}, {1, 0}, {20, 0}, {3, 0}, {4, 0}}
	expected := 20.0

	solution := solve(input)

	result := big.NewFloat(solution).Cmp(big.NewFloat(expected))

	if 0 != result {
		t.Errorf("Most distant edge case x was incorrect, got: %f, want: %f.", solution, expected)
	}
}

// TestMostDistantEdgeYCase implements the test given as edge y example on hackerrank.
func TestMostDistantEdgeYCase(t *testing.T) {
	input := [][]int64{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 9}}
	expected := 9.0

	solution := solve(input)

	result := big.NewFloat(solution).Cmp(big.NewFloat(expected))

	if 0 != result {
		t.Errorf("Most distant edge y case was incorrect, got: %f, want: %f.", solution, expected)
	}
}

// TestMostDistantEdgeCase implements the test given as edge example on hackerrank.
func TestMostDistantEdgeCase(t *testing.T) {
	input := [][]int64{{2, 0}, {0, 2}, {-2, 0}, {0, -2}, {9, 0}}
	expected := 11.0

	solution := solve(input)

	result := big.NewFloat(solution).Cmp(big.NewFloat(expected))

	if 0 != result {
		t.Errorf("Most distant edge case was incorrect, got: %f, want: %f.", solution, expected)
	}
}
