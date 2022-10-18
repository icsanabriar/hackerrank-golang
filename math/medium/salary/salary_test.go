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

// TestSalaryBluesGivenCase implements the test given as first example on hackerrank.
func TestSalaryBluesGivenCase(t *testing.T) {

	a := []int64{9, 12, 3, 6}
	queries := []int32{0, 3}
	expected := []int64{3, 3}

	result := normalize(a, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Salary blues first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestSalaryBluesEdgeGivenCase implements the test given as edge example on hackerrank.
func TestSalaryBluesEdgeGivenCase(t *testing.T) {

	a := []int64{474515580, 11940082770, 20257709010, 20257709010}
	queries := []int32{0, 868972117, 971690506, 877902860}
	expected := []int64{2310, 1, 2, 10}

	result := normalize(a, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Salary blues edge case was incorrect, got: %d, want: %d.", result, expected)
	}
}
