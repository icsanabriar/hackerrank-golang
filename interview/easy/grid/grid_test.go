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

import "testing"

// TestGridChallengeFirstGivenCase implements the test given as first example on hackerrank.
func TestGridChallengeFirstGivenCase(t *testing.T) {
	grid := []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}
	expected := "YES"

	result := gridChallenge(grid)

	if result != expected {
		t.Errorf("Grid challenge first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestGridChallengeSecondGivenCase implements the test given as second example on hackerrank.
func TestGridChallengeSecondGivenCase(t *testing.T) {
	grid := []string{"zb", "ac"}
	expected := "NO"

	result := gridChallenge(grid)

	if result != expected {
		t.Errorf("Grid challenge second case was incorrect, got: %s, want: %s.", result, expected)
	}
}
