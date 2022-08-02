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

// TestRotationFirstGivenCase implements the test given as first example on hackerrank.
func TestRotationFirstGivenCase(t *testing.T) {

	a := []int32{1, 2, 3, 4, 5}
	d := int32(4)

	expected := []int32{5, 1, 2, 3, 4}

	result := rotLeft(a, d)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Left Rotation first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestRotationSecondGivenCase implements the test given as second example on hackerrank.
func TestRotationSecondGivenCase(t *testing.T) {

	a := []int32{1, 2, 3, 4, 5}
	d := int32(2)

	expected := []int32{3, 4, 5, 1, 2}

	result := rotLeft(a, d)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Left Rotation second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
