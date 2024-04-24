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

// TestSalesByMatchFirstGivenCase implements the test given as first example on hackerrank.
func TestSalesByMatchFirstGivenCase(t *testing.T) {
	n := 7
	input := []int64{1, 2, 1, 2, 1, 3, 2}
	expected := int64(2)

	result := sockMerchant(n, input)

	if result != expected {
		t.Errorf("Sales by match first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestSalesByMatchSecondGivenCase implements the test given as second example on hackerrank.
func TestSalesByMatchSecondGivenCase(t *testing.T) {
	n := 9
	input := []int64{10, 20, 20, 10, 10, 30, 50, 10, 20}
	expected := int64(3)

	result := sockMerchant(n, input)

	if result != expected {
		t.Errorf("Sales by match second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
