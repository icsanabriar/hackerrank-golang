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

// TestSparseFirstGivenCase implements the test given as first example on hackerrank.
func TestSparseFirstGivenCase(t *testing.T) {
	strings := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}

	expected := []int32{2, 1, 0}

	result := matchingStrings(strings, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sparse arrays first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestSparseSecondGivenCase implements the test given as second example on hackerrank.
func TestSparseSecondGivenCase(t *testing.T) {
	strings := []string{"def", "de", "fgh"}
	queries := []string{"de", "lmn", "fgh"}

	expected := []int32{1, 0, 1}

	result := matchingStrings(strings, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sparse arrays second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
