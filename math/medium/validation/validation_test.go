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

// TestIsFiboFirstGivenCase implements the test given as first example on hackerrank.
func TestIsFiboFirstGivenCase(t *testing.T) {

	cache := buildCache()

	input := int64(5)
	expected := "IsFibo"

	result := solveIsFibo(input, cache)

	if result != expected {
		t.Errorf("IsFibo first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestIsFiboSecondGivenCase implements the test given as second example on hackerrank.
func TestIsFiboSecondGivenCase(t *testing.T) {

	cache := buildCache()

	input := int64(7)
	expected := "IsNotFibo"

	result := solveIsFibo(input, cache)

	if result != expected {
		t.Errorf("IsFibo second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestIsFiboThirdGivenCase implements the test given as third example on hackerrank.
func TestIsFiboThirdGivenCase(t *testing.T) {

	cache := buildCache()

	input := int64(8)
	expected := "IsFibo"

	result := solveIsFibo(input, cache)

	if result != expected {
		t.Errorf("IsFibo third case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// buildCache factory of a valid cache for solveIsFibo function.
func buildCache() []int64 {

	cache := make([]int64, 100)

	for i := range cache {
		cache[i] = -1
	}

	return cache
}
