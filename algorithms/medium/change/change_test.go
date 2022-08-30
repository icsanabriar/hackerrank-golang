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

// TestCoinChangeProblemFirstGivenCase implements the test given as first example on hackerrank.
func TestCoinChangeProblemFirstGivenCase(t *testing.T) {

	n := int64(4)
	c := []int64{1, 2, 3}
	expected := int64(4)

	result := getWays(n, c)

	if result != expected {
		t.Errorf("Coin change problem first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCoinChangeProblemSecondGivenCase implements the test given as second example on hackerrank.
func TestCoinChangeProblemSecondGivenCase(t *testing.T) {

	n := int64(10)
	c := []int64{2, 5, 3, 6}
	expected := int64(5)

	result := getWays(n, c)

	if result != expected {
		t.Errorf("Coin change problem second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCoinChangeProblemHiddenCase implements the test given as hidden example on hackerrank.
func TestCoinChangeProblemHiddenCase(t *testing.T) {

	n := int64(219)
	c := []int64{36, 10, 42, 7, 50, 1, 49, 24, 37, 12, 34, 13, 39, 18, 8, 29, 19, 43, 5, 44, 28, 23, 35, 26}
	expected := int64(168312708)

	result := getWays(n, c)

	if result != expected {
		t.Errorf("Coin change problem hidden case was incorrect, got: %d, want: %d.", result, expected)
	}
}
