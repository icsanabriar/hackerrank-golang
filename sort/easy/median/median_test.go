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

// TestFindingMedianGivenCase implements the test given as first example on hackerrank.
func TestFindingMedianGivenCase(t *testing.T) {
	arr := []int64{0, 1, 2, 4, 6, 5, 3}
	expected := int64(3)

	result := findMedian(arr)

	if result != expected {
		t.Errorf("Finding median first case was incorrect, got: %d, want: %d.", result, expected)
	}
}
