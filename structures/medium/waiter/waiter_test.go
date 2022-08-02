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

// TestWaiterFirstGivenCase implements the test given as first example on hackerrank.
func TestWaiterFirstGivenCase(t *testing.T) {

	number := []int32{3, 4, 7, 6, 5}
	q := 1

	expected := []int32{4, 6, 3, 7, 5}

	result := waiter(number, q)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Waiter first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestWaiterSecondGivenCase implements the test given as second example on hackerrank.
func TestWaiterSecondGivenCase(t *testing.T) {

	number := []int32{3, 3, 4, 4, 9}
	q := 2

	expected := []int32{4, 4, 9, 3, 3}

	result := waiter(number, q)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Waiter second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
