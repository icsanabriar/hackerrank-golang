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

// TestConnectingTownsFirstGivenCase implements the test given as first example on hackerrank.
func TestConnectingTownsFirstGivenCase(t *testing.T) {

	n := int32(3)
	routes := []int32{1, 3}
	expected := int32(3)

	result := connectingTowns(n, routes)

	if result != expected {
		t.Errorf("Connecting Towns first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestConnectingTownsSecondGivenCase implements the test given as second example on hackerrank.
func TestConnectingTownsSecondGivenCase(t *testing.T) {

	n := int32(4)
	routes := []int32{2, 2, 2}
	expected := int32(8)

	result := connectingTowns(n, routes)

	if result != expected {
		t.Errorf("Connecting Towns second case was incorrect, got: %d, want: %d.", result, expected)
	}
}