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

import "testing"

// TestRestaurantFirstGivenCase implements the test given as first example on hackerrank.
func TestRestaurantFirstGivenCase(t *testing.T) {
	l := int64(2)
	b := int64(2)
	expected := int64(1)

	result := restaurant(l, b)

	if result != expected {
		t.Errorf("Restaurant first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestRestaurantSecondGivenCase implements the test given as second example on hackerrank.
func TestRestaurantSecondGivenCase(t *testing.T) {
	l := int64(6)
	b := int64(9)
	expected := int64(6)

	result := restaurant(l, b)

	if result != expected {
		t.Errorf("Restaurant second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
