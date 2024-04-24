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

// TestTowerBreakersFirstGivenCase implements the test given as first example on hackerrank.
func TestTowerBreakersFirstGivenCase(t *testing.T) {
	n := int64(2)
	m := int64(2)
	expected := int64(2)

	result := towerBreakers(n, m)

	if result != expected {
		t.Errorf("Tower breakers first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestTowerBreakersSecondGivenCase implements the test given as second example on hackerrank.
func TestTowerBreakersSecondGivenCase(t *testing.T) {
	n := int64(1)
	m := int64(4)
	expected := int64(1)

	result := towerBreakers(n, m)

	if result != expected {
		t.Errorf("Tower breakers first case was incorrect, got: %d, want: %d.", result, expected)
	}
}
