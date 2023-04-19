// Copyright 2021 Iván Camilo Sanabria
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

// TestHandshakeFirstGivenCase implements the test given as first example on hackerrank.
func TestHandshakeFirstGivenCase(t *testing.T) {

	input := int32(1)
	expected := int32(0)

	result := handshake(input)

	if result != expected {
		t.Errorf("Handshake first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestHandshakeSecondGivenCase implements the test given as second example on hackerrank.
func TestHandshakeSecondGivenCase(t *testing.T) {

	input := int32(2)
	expected := int32(1)

	result := handshake(input)

	if result != expected {
		t.Errorf("Handshake second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestHandshakeEdgeCase implements the edge case to validate results on hackerrank.
func TestHandshakeEdgeCase(t *testing.T) {

	input := int32(100)
	expected := int32(4950)

	result := handshake(input)

	if result != expected {
		t.Errorf("Handshake edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
