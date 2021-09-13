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

// TestJesseAndCookiesFirstGivenCase implements the test given as first example on hackerrank.
func TestJesseAndCookiesFirstGivenCase(t *testing.T) {

	k := int64(9)
	A := []int32{2, 7, 3, 6, 4, 6}

	expected := int64(4)

	result := cookies(k, A)

	if result != expected {
		t.Errorf("Jesse and Cookies first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestJesseAndCookiesSecondGivenCase implements the test given as second example on hackerrank.
func TestJesseAndCookiesSecondGivenCase(t *testing.T) {

	k := int64(7)
	A := []int32{1, 2, 3, 9, 10, 12}

	expected := int64(2)

	result := cookies(k, A)

	if result != expected {
		t.Errorf("Jesse and Cookies second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestJesseAndCookiesEdgeGivenCase implements the test given as edge example on hackerrank.
func TestJesseAndCookiesEdgeGivenCase(t *testing.T) {

	k := int64(10)
	A := []int32{1, 1, 1}

	expected := int64(-1)

	result := cookies(k, A)

	if result != expected {
		t.Errorf("Jesse and Cookies edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
