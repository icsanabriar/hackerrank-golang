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

// TestDigitsFirstGivenCase implements the test given as first example on hackerrank.
func TestDigitsFirstGivenCase(t *testing.T) {

	input := int32(12)
	expected := 2

	result := findDigits(input)

	if result != expected {
		t.Errorf("Finding digits first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestDigitsSecondGivenCase implements the test given as second example on hackerrank.
func TestDigitsSecondGivenCase(t *testing.T) {

	input := int32(1012)
	expected := 3

	result := findDigits(input)

	if result != expected {
		t.Errorf("Finding digits second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
