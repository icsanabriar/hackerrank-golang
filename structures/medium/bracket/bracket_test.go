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

// TestBalanceBracketFirstGivenCase implements the test given as first example on hackerrank.
func TestBalanceBracketFirstGivenCase(t *testing.T) {

	s := "{[()]}"
	expected := "YES"

	result := isBalanced(s)

	if result != expected {
		t.Errorf("Balance bracket first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestBalanceBracketSecondGivenCase implements the test given as second example on hackerrank.
func TestBalanceBracketSecondGivenCase(t *testing.T) {

	s := "{[(])}"
	expected := "NO"

	result := isBalanced(s)

	if result != expected {
		t.Errorf("Balance bracket second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestBalanceBracketThirdGivenCase implements the test given as third example on hackerrank.
func TestBalanceBracketThirdGivenCase(t *testing.T) {

	s := "{{[[(())]]}}"
	expected := "YES"

	result := isBalanced(s)

	if result != expected {
		t.Errorf("Balance bracket third case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestBalanceBracketHiddenCase implements the test given as hidden example on hackerrank.
func TestBalanceBracketHiddenCase(t *testing.T) {

	s := "}([[{)[]))]{){}["
	expected := "NO"

	result := isBalanced(s)

	if result != expected {
		t.Errorf("Balance bracket hidden case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestBalanceBracketEdgeCase implements the test given as edge example.
func TestBalanceBracketEdgeCase(t *testing.T) {

	s := "{[()]"
	expected := "NO"

	result := isBalanced(s)

	if result != expected {
		t.Errorf("Balance bracket edge case was incorrect, got: %s, want: %s.", result, expected)
	}
}
