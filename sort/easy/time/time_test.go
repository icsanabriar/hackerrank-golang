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

// TestRunningTimeGivenCase implements the test given as first example on hackerrank.
func TestRunningTimeGivenCase(t *testing.T) {
	arr := []int64{2, 1, 3, 1, 2}
	expected := int64(4)

	result := runningTime(arr)

	if result != expected {
		t.Errorf("Running time first case was incorrect, got: %d, want: %d.", result, expected)
	}
}
