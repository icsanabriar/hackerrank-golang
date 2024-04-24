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

// TestAbsolutePermutationFirstGivenCase implements the test given as first example on hackerrank.
func TestAbsolutePermutationFirstGivenCase(t *testing.T) {
	n := int64(4)
	k := int64(2)
	expected := []int64{3, 4, 1, 2}

	result := absolutePermutation(n, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Absolute permutation first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestAbsolutePermutationSecondGivenCase implements the test given as second example on hackerrank.
func TestAbsolutePermutationSecondGivenCase(t *testing.T) {
	n := int64(2)
	k := int64(1)
	expected := []int64{2, 1}

	result := absolutePermutation(n, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Absolute permutation second case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestAbsolutePermutationThirdGivenCase implements the test given as third example on hackerrank.
func TestAbsolutePermutationThirdGivenCase(t *testing.T) {
	n := int64(3)
	k := int64(0)
	expected := []int64{1, 2, 3}

	result := absolutePermutation(n, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Absolute permutation third case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestAbsolutePermutationFourthGivenCase implements the test given as fourth example on hackerrank.
func TestAbsolutePermutationFourthGivenCase(t *testing.T) {
	n := int64(3)
	k := int64(2)
	expected := []int64{-1}

	result := absolutePermutation(n, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Absolute permutation fourth case was incorrect, got: %v, want: %v.", result, expected)
	}
}
