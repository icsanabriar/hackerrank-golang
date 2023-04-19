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

// TestDieHard3FirstGivenCase implements the test given as first example on hackerrank.
func TestDieHard3FirstGivenCase(t *testing.T) {

	a := int32(5)
	b := int32(3)
	c := int32(4)
	expected := "YES"

	result := solve(a, b, c)

	if result != expected {
		t.Errorf("Die hard 3 first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestDieHard3SecondGivenCase implements the test given as second example on hackerrank.
func TestDieHard3SecondGivenCase(t *testing.T) {

	a := int32(3)
	b := int32(6)
	c := int32(4)
	expected := "NO"

	result := solve(a, b, c)

	if result != expected {
		t.Errorf("Die hard 3 second case was incorrect, got: %s, want: %s.", result, expected)
	}
}
