// Copyright 2021 Iván Camilo Sanabria
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

// TestMutualRecurrenceFirstGivenCase implements the test given as first example on hackerrank.
func TestMutualRecurrenceFirstGivenCase(t *testing.T) {

	input := []int32{1, 2, 3, 1, 1, 2, 3, 1}
	n := int64(10)
	expected := []int64{1910, 1910}

	result := recurrence(input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7], n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Mutual recurrence first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestMutualRecurrenceSecondGivenCase implements the test given as second example on hackerrank.
func TestMutualRecurrenceSecondGivenCase(t *testing.T) {

	input := []int32{1, 2, 3, 2, 2, 1, 1, 4}
	n := int64(10)
	expected := []int64{909323, 11461521}

	result := recurrence(input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7], n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Mutual recurrence second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestMutualRecurrenceThirdGivenCase implements the test given as third example on hackerrank.
func TestMutualRecurrenceThirdGivenCase(t *testing.T) {

	input := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	n := int64(90)
	expected := []int64{108676813, 414467031}

	result := recurrence(input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7], n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Mutual recurrence third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
