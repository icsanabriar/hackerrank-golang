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

// TestCandiesFirstGivenCase implements the test given as first example on hackerrank.
func TestCandiesFirstGivenCase(t *testing.T) {
	n := int64(10)
	arr := []int64{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}
	expected := int64(19)

	result := candies(n, arr)

	if result != expected {
		t.Errorf("Candies first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestCandiesSecondGivenCase implements the test given as second example on hackerrank.
func TestCandiesSecondGivenCase(t *testing.T) {
	n := int64(8)
	arr := []int64{2, 4, 3, 5, 2, 6, 4, 5}
	expected := int64(12)

	result := candies(n, arr)

	if result != expected {
		t.Errorf("Candies second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
