// Copyright 2022 Iván Camilo Sanabria
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

// TestJimJokesFirstGivenCase implements the test given as first example on hackerrank.
func TestJimJokesFirstGivenCase(t *testing.T) {

	a := [][]int32{{10, 25}, {8, 31}}
	expected := int64(1)

	result := solve(a)

	if result != expected {
		t.Errorf("Jim and the Jokes first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestJimJokesSecondGivenCase implements the test given as second example on hackerrank.
func TestJimJokesSecondGivenCase(t *testing.T) {

	a := [][]int32{{2, 25}, {2, 25}}
	expected := int64(0)

	result := solve(a)

	if result != expected {
		t.Errorf("Jim and the Jokes second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestJimJokesThirdGivenCase implements the test given as third example on hackerrank.
func TestJimJokesThirdGivenCase(t *testing.T) {

	a := [][]int32{{11, 10}, {10, 11}}
	expected := int64(1)

	result := solve(a)

	if result != expected {
		t.Errorf("Jim and the Jokes third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
