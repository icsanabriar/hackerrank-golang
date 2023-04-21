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

// TestPalindromeIndexFirstGivenCase implements the test given as first example on hackerrank.
func TestPalindromeIndexFirstGivenCase(t *testing.T) {

	s := "aaab"
	expected := 3

	result := palindromeIndex(s)

	if result != expected {
		t.Errorf("Palindrome index first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestPalindromeIndexSecondGivenCase implements the test given as second example on hackerrank.
func TestPalindromeIndexSecondGivenCase(t *testing.T) {

	s := "baa"
	expected := 0

	result := palindromeIndex(s)

	if result != expected {
		t.Errorf("Palindrome index second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestPalindromeIndexThirdGivenCase implements the test given as third example on hackerrank.
func TestPalindromeIndexThirdGivenCase(t *testing.T) {

	s := "aaa"
	expected := -1

	result := palindromeIndex(s)

	if result != expected {
		t.Errorf("Palindrome index third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestPalindromeIndexEdgeCase implements the test as edge example.
func TestPalindromeIndexEdgeCase(t *testing.T) {

	s := "arcade"
	expected := -1

	result := palindromeIndex(s)

	if result != expected {
		t.Errorf("Palindrome index edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestPalindromeIndexEmptyCase implements the test as empty input example.
func TestPalindromeIndexEmptyCase(t *testing.T) {

	s := ""
	expected := -1

	result := palindromeIndex(s)

	if result != expected {
		t.Errorf("Palindrome index empty case was incorrect, got: %d, want: %d.", result, expected)
	}
}
