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

// TestLegoBlocksFirstGivenCase implements the test given as first example on hackerrank.
func TestLegoBlocksFirstGivenCase(t *testing.T) {
	n := 3
	m := 2
	expected := int64(7)

	result := legoBlocks(n, m)

	if result != expected {
		t.Errorf("Lego blocks first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLegoBlocksSecondGivenCase implements the test given as second example on hackerrank.
func TestLegoBlocksSecondGivenCase(t *testing.T) {
	n := 2
	m := 3
	expected := int64(9)

	result := legoBlocks(n, m)

	if result != expected {
		t.Errorf("Lego blocks second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLegoBlocksThirdGivenCase implements the test given as third example on hackerrank.
func TestLegoBlocksThirdGivenCase(t *testing.T) {
	n := 4
	m := 4
	expected := int64(3375)

	result := legoBlocks(n, m)

	if result != expected {
		t.Errorf("Lego blocks third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLegoBlocksHiddenCase implements the test given as hidden example on hackerrank.
func TestLegoBlocksHiddenCase(t *testing.T) {
	n := 924
	m := 604
	expected := int64(382238489)

	result := legoBlocks(n, m)

	if result != expected {
		t.Errorf("Lego blocks third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLegoBlocksBasicCase implements the test for basic case.
func TestLegoBlocksBasicCase(t *testing.T) {
	n := 2
	m := 1
	expected := int64(1)

	result := legoBlocks(n, m)

	if result != expected {
		t.Errorf("Lego blocks basic case was incorrect, got: %d, want: %d.", result, expected)
	}
}
