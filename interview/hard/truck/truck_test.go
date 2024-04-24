// Copyright 2023 Iván Camilo Sanabria
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

// TestTruckTourFirstGivenCase implements the test given as first example on hackerrank.
func TestTruckTourFirstGivenCase(t *testing.T) {
	pump := [][]int64{{1, 5}, {10, 3}, {3, 4}}
	expected := int64(1)

	result := truckTour(pump)

	if result != expected {
		t.Errorf("Truck tour first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestTruckTourSecondGivenCase implements the test given as second example on hackerrank.
func TestTruckTourSecondGivenCase(t *testing.T) {
	pump := [][]int64{{2, 3}, {2, 2}, {1, 10}}
	expected := int64(-1)

	result := truckTour(pump)

	if result != expected {
		t.Errorf("Truck tour second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestTruckTourEdgeCase implements the test given as edge example.
func TestTruckTourEdgeCase(t *testing.T) {
	pump := [][]int64{{6, 3}, {1, 4}, {1, 1}}
	expected := int64(-1)

	result := truckTour(pump)

	if result != expected {
		t.Errorf("Truck tour edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
