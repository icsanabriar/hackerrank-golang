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

// TestNoDivisibleSubsetFirstGivenCase implements the test given as first example on hackerrank.
func TestNoDivisibleSubsetFirstGivenCase(t *testing.T) {

	k := int32(3)
	s := []int32{1, 7, 2, 4}
	expected := int32(3)

	result := nonDivisibleSubset(k, s)

	if result != expected {
		t.Errorf("No divisible subset first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestNoDivisibleSubsetSecondGivenCase implements the test given as second example on hackerrank.
func TestNoDivisibleSubsetSecondGivenCase(t *testing.T) {

	k := int32(4)
	s := []int32{19, 10, 12, 10, 24, 25, 22}
	expected := int32(3)

	result := nonDivisibleSubset(k, s)

	if result != expected {
		t.Errorf("No divisible subset second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestNoDivisibleSubsetThirdGivenCase implements the test given as third example on hackerrank.
func TestNoDivisibleSubsetThirdGivenCase(t *testing.T) {

	k := int32(7)
	s := []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}
	expected := int32(11)

	result := nonDivisibleSubset(k, s)

	if result != expected {
		t.Errorf("No divisible subset third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
