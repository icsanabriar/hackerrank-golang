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
	"reflect"
	"testing"
)

// TestQuickSortGivenCase implements the test given as first example on hackerrank.
func TestQuickSortGivenCase(t *testing.T) {

	arr := []int32{4, 5, 3, 7, 2}
	expected := []int32{3, 2, 4, 5, 7}

	result := quickSort(arr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Quick sort first case was incorrect, got: %v, want: %v.", result, expected)
	}
}
