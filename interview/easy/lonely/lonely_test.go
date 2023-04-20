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

import (
	"testing"
)

// TestLonelyIntegerFirstGivenCase implements the test given as first example on hackerrank.
func TestLonelyIntegerFirstGivenCase(t *testing.T) {

	arr := []int32{1, 2, 3, 4, 3, 2, 1}
	expected := int32(4)

	result := lonelyInteger(arr)

	if result != expected {
		t.Errorf("Lonely integer first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLonelyIntegerEdgeCase implements the test given as edge example.
func TestLonelyIntegerEdgeCase(t *testing.T) {

	arr := []int32{1, 2, 3, 4, 3, 2, 1, 4}
	expected := int32(0)

	result := lonelyInteger(arr)

	if result != expected {
		t.Errorf("Lonely integer edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
