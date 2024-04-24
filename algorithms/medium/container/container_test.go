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

// TestOrganizingContainerFirstGivenCase implements the test given as first example on hackerrank.
func TestOrganizingContainerFirstGivenCase(t *testing.T) {
	size := 2
	container := [][]int64{{1, 1}, {1, 1}}
	expected := "Possible"

	result := organizingContainers(container, size)

	if result != expected {
		t.Errorf("Organizing container first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestOrganizingContainerSecondGivenCase implements the test given as second example on hackerrank.
func TestOrganizingContainerSecondGivenCase(t *testing.T) {
	size := 2
	container := [][]int64{{0, 2}, {1, 1}}
	expected := "Impossible"

	result := organizingContainers(container, size)

	if result != expected {
		t.Errorf("Organizing container second case was incorrect, got: %s, want: %s.", result, expected)
	}
}
