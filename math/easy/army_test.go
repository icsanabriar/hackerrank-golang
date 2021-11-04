// Copyright 2021 Iván Camilo Sanabria
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

// TestArmyGameGivenCase implements the test given as first example on hackerrank.
func TestArmyGameGivenCase(t *testing.T) {

	n := int32(3)
	m := int32(2)
	expected := int32(2)

	result := gameWithCells(n, m)

	if result != expected {
		t.Errorf("Army game first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestArmyGameHiddenCase implements the test given as hidden example on hackerrank.
func TestArmyGameHiddenCase(t *testing.T) {

	n := int32(39)
	m := int32(169)
	expected := int32(1700)

	result := gameWithCells(n, m)

	if result != expected {
		t.Errorf("Army game hidden case was incorrect, got: %d, want: %d.", result, expected)
	}
}
