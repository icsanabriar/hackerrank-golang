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

import (
	"reflect"
	"testing"
)

// TestDynamicArrayFirstGivenCase implements the test given as first example on hackerrank.
func TestDynamicArrayFirstGivenCase(t *testing.T) {

	queries := [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}}
	n := int32(2)

	expected := []int32{7, 3}

	result := dynamicArray(n, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Dynamic Array first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestDynamicArraySecondGivenCase implements the test given as second example on hackerrank.
func TestDynamicArraySecondGivenCase(t *testing.T) {

	queries := [][]int32{{1, 0, 3}, {1, 1, 1}, {1, 0, 0}, {2, 2, 2}}
	n := int32(2)

	expected := []int32{3}

	result := dynamicArray(n, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Dynamic Array second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
