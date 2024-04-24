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

// TestCommonChildFirstGivenCase implements the test given as first example on hackerrank.
func TestCommonChildFirstGivenCase(t *testing.T) {
	s1 := "ABCD"
	s2 := "ABDC"
	expected := int32(3)

	result := commonChild(s1, s2)

	if result != expected {
		t.Errorf("Common child first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCommonChildSecondGivenCase implements the test given as second example on hackerrank.
func TestCommonChildSecondGivenCase(t *testing.T) {
	s1 := "HARRY"
	s2 := "SALLY"
	expected := int32(2)

	result := commonChild(s1, s2)

	if result != expected {
		t.Errorf("Common child second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestCommonChildThirdGivenCase implements the test given as third example on hackerrank.
func TestCommonChildThirdGivenCase(t *testing.T) {
	s1 := "AA"
	s2 := "BB"
	expected := int32(0)

	result := commonChild(s1, s2)

	if result != expected {
		t.Errorf("Common child third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
