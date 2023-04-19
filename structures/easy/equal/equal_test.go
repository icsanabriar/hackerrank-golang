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

import "testing"

// TestEqualStackFirstGivenCase implements the test given as first example on hackerrank.
func TestEqualStackFirstGivenCase(t *testing.T) {

	h1 := []int32{1, 2, 1, 1}
	h2 := []int32{1, 1, 2}
	h3 := []int32{1, 1}

	expected := int64(2)

	result := equalStacks(h1, h2, h3)

	if result != expected {
		t.Errorf("Equal stack first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestEqualStackSecondGivenCase implements the test given as second example on hackerrank.
func TestEqualStackSecondGivenCase(t *testing.T) {

	h1 := []int32{3, 2, 1, 1, 1}
	h2 := []int32{4, 3, 2}
	h3 := []int32{1, 1, 4, 1}

	expected := int64(5)

	result := equalStacks(h1, h2, h3)

	if result != expected {
		t.Errorf("Equal stack second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestEqualStackThirdGivenCase implements the test given as third example on hackerrank.
func TestEqualStackThirdGivenCase(t *testing.T) {

	h1 := []int32{1, 1, 1, 1, 2}
	h2 := []int32{3, 7}
	h3 := []int32{1, 3, 1}

	expected := int64(0)

	result := equalStacks(h1, h2, h3)

	if result != expected {
		t.Errorf("Equal stack third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
