// Copyright 2020 Iván Camilo Sanabria
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

// TestWordsFirstGivenCase implements the test given as first example on hackerrank.
func TestWordsFirstGivenCase(t *testing.T) {

	h := int32(5)
	m := int32(47)
	expected := "thirteen minutes to six"

	result := timeInWords(h, m)

	if result != expected {
		t.Errorf("The time in words first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestWordsSecondGivenCase implements the test given as second example on hackerrank.
func TestWordsSecondGivenCase(t *testing.T) {

	h := int32(3)
	m := int32(0)
	expected := "three o' clock"

	result := timeInWords(h, m)

	if result != expected {
		t.Errorf("The time in words second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestWordsThirdGivenCase implements the test given as third example on hackerrank.
func TestWordsThirdGivenCase(t *testing.T) {

	h := int32(1)
	m := int32(1)
	expected := "one minute past one"

	result := timeInWords(h, m)

	if result != expected {
		t.Errorf("The time in words third case was incorrect, got: %s, want: %s.", result, expected)
	}
}
