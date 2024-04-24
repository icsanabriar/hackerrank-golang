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

// TestSticksFirstGivenCase implements the test given as first example on hackerrank.
func TestSticksFirstGivenCase(t *testing.T) {
	input := []int64{5, 4, 4, 2, 2, 8}
	expected := []int64{6, 4, 2, 1}

	result := cutTheSticks(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Cut the sticks first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestSticksSecondGivenCase implements the test given as second example on hackerrank.
func TestSticksSecondGivenCase(t *testing.T) {
	input := []int64{1, 2, 3, 4, 3, 3, 2, 1}
	expected := []int64{8, 6, 4, 1}

	result := cutTheSticks(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Cut the sticks second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
