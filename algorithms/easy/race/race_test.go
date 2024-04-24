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

// TestRaceFirstGivenCase implements the test given as first example on hackerrank.
func TestRaceFirstGivenCase(t *testing.T) {
	k := int64(4)
	h := []int64{1, 6, 3, 5, 2}
	expected := int64(2)

	result := hurdleRace(k, h)

	if result != expected {
		t.Errorf("The Hurdle race first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRaceSecondGivenCase implements the test given as second example on hackerrank.
func TestRaceSecondGivenCase(t *testing.T) {
	k := int64(7)
	h := []int64{2, 5, 4, 5, 2}
	expected := int64(0)

	result := hurdleRace(k, h)

	if result != expected {
		t.Errorf("The Hurdle race second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
