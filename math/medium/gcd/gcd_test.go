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

import (
	"testing"
)

// TestEasyGcdFirstGivenCase implements the test given as first example on hackerrank.
func TestEasyGcdFirstGivenCase(t *testing.T) {
	k := int64(5)
	arr := []int64{2, 6, 4}
	expected := int64(4)

	result := easyGcd(k, arr)

	if result != expected {
		t.Errorf("Easy gcd first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestEasyGcdSecondGivenCase implements the test given as second example on hackerrank.
func TestEasyGcdSecondGivenCase(t *testing.T) {
	k := int64(5)
	arr := []int64{7}
	expected := int64(0)

	result := easyGcd(k, arr)

	if result != expected {
		t.Errorf("Easy gcd second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestEasyGcdHiddenCase implements the test given as hidden example on hackerrank.
func TestEasyGcdHiddenCase(t *testing.T) {
	k := int64(999999999)
	arr := []int64{1000000000, 1000000000}
	expected := int64(999999998)

	result := easyGcd(k, arr)

	if result != expected {
		t.Errorf("Easy gcd hidden case was incorrect, got: %d, want: %d.", result, expected)
	}
}
