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

// TestPairsFirstGivenCase implements the test given as first example on hackerrank.
func TestPairsFirstGivenCase(t *testing.T) {
	k := int64(1)
	arr := []int64{1, 2, 3, 4}
	expected := int64(3)

	result := pairs(k, arr)

	if result != expected {
		t.Errorf("Pairs first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestPairsSecondGivenCase implements the test given as second example on hackerrank.
func TestPairsSecondGivenCase(t *testing.T) {
	k := int64(2)
	arr := []int64{1, 5, 3, 4, 2}
	expected := int64(3)

	result := pairs(k, arr)

	if result != expected {
		t.Errorf("Pairs second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
