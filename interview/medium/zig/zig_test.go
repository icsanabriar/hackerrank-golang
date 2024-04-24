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
	"reflect"
	"testing"
)

// TestZigZagSequenceFirstGivenCase implements the test given as first example on hackerrank.
func TestZigZagSequenceFirstGivenCase(t *testing.T) {
	arr := []int64{2, 3, 5, 1, 4}
	expected := []int64{1, 2, 5, 4, 3}

	result := zigZagSequence(arr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zig zag sequence first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestZigZagSequenceSecondGivenCase implements the test given as second example on hackerrank.
func TestZigZagSequenceSecondGivenCase(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7}
	expected := []int64{1, 2, 3, 7, 6, 5, 4}

	result := zigZagSequence(arr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zig zag sequence second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
