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

import (
	"reflect"
	"testing"
)

// TestContactsFirstGivenCase implements the test given as first example on hackerrank.
func TestContactsFirstGivenCase(t *testing.T) {

	input := [][]string{{"add", "hack"},
		{"add", "hackerrank"},
		{"find", "hac"},
		{"find", "hak"}}

	expected := []int32{2, 0}

	result := contacts(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Contacts first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestContactsSecondGivenCase implements the test given as second example on hackerrank.
func TestContactsSecondGivenCase(t *testing.T) {

	input := [][]string{{"add", "s"},
		{"add", "ss"},
		{"add", "sss"},
		{"add", "ssss"},
		{"add", "sssss"},
		{"find", "s"},
		{"find", "ss"},
		{"find", "sss"},
		{"find", "ssss"},
		{"find", "sssss"},
		{"find", "ssssss"}}

	expected := []int32{5, 4, 3, 2, 1, 0}

	result := contacts(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Contacts second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
