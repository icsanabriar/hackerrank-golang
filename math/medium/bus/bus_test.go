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

// TestBusStationGivenCase implements the test given as first example on hackerrank.
func TestBusStationGivenCase(t *testing.T) {
	a := []int64{1, 2, 1, 1, 1, 2, 1, 3}
	expected := []int64{3, 4, 6, 12}

	result := solve(a)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Bus station first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestBusStationHiddenCase implements the test given as hidden example on hackerrank.
func TestBusStationHiddenCase(t *testing.T) {
	a := []int64{1, 1, 1, 1}
	expected := []int64{1, 2, 4}

	result := solve(a)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Bus station hidden case was incorrect, got: %v, want: %v.", result, expected)
	}
}
