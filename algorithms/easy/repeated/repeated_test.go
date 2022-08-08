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

// TestRepeatedFirstGivenCase implements the test given as first example on hackerrank.
func TestRepeatedFirstGivenCase(t *testing.T) {

	input := "aba"
	n := int64(10)

	expected := int64(7)

	result := repeatedString(input, n)

	if result != expected {
		t.Errorf("Repeated string first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRepeatedSecondGivenCase implements the test given as second example on hackerrank.
func TestRepeatedSecondGivenCase(t *testing.T) {

	input := "a"
	n := int64(1000000000000)

	expected := int64(1000000000000)

	result := repeatedString(input, n)

	if result != expected {
		t.Errorf("Repeated string second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRepeatedThirdGivenCase implements the test given as third example on hackerrank.
func TestRepeatedThirdGivenCase(t *testing.T) {

	input := "epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq"
	n := int64(549382313570)

	expected := int64(16481469408)

	result := repeatedString(input, n)

	if result != expected {
		t.Errorf("Repeated string third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
