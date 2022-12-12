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

import "testing"

// TestCloudyDayGivenCase implements the test given as first example on hackerrank.
func TestCloudyDayGivenCase(t *testing.T) {

	p := []int64{10, 100}
	x := []int64{5, 100}
	y := []int64{4}
	r := []int64{1}
	expected := int64(110)

	result := maximumPeople(p, x, y, r)

	if result != expected {
		t.Errorf("Cloudy day first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestCloudyDayHiddenCase implements the test given as hidden example on hackerrank.
func TestCloudyDayHiddenCase(t *testing.T) {

	p := []int64{10, 1, 8, 3}
	x := []int64{4, 5, 7, 2}
	y := []int64{3, 9, 3, 5}
	r := []int64{11, 10, 8, 7}
	expected := int64(0)

	result := maximumPeople(p, x, y, r)

	if result != expected {
		t.Errorf("Cloudy day hidden case was incorrect, got: %v, want: %v.", result, expected)
	}
}
