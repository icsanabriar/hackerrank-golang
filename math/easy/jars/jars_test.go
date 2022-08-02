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

import "testing"

// TestFillingJarsGivenCase implements the test given as first example on hackerrank.
func TestFillingJarsGivenCase(t *testing.T) {

	n := int64(5)
	operations := [][]int64{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}
	expected := int64(160)

	result := estimate(n, operations)

	if result != expected {
		t.Errorf("Filling jars first case was incorrect, got: %d, want: %d.", result, expected)
	}
}
