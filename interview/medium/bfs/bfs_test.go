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

import (
	"reflect"
	"testing"
)

// TestBFSFirstGivenCase implements the test given as first example on hackerrank.
func TestBFSFirstGivenCase(t *testing.T) {

	n := int32(4)
	m := int32(2)
	edges := [][]int32{{1, 2}, {1, 3}}
	s := int32(1)

	expected := []int32{6, 6, -1}

	result := bfs(n, m, edges, s)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestBFSSecondGivenCase implements the test given as second example on hackerrank.
func TestBFSSecondGivenCase(t *testing.T) {

	n := int32(3)
	m := int32(1)
	edges := [][]int32{{2, 3}}
	s := int32(2)

	expected := []int32{-1, 6}

	result := bfs(n, m, edges, s)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS second case was incorrect, got: %d, want: %d.", result, expected)
	}
}
