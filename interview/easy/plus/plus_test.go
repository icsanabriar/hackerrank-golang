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

// TestPlusMinusFirstGivenCase implements the test given as first example on hackerrank.
func TestPlusMinusFirstGivenCase(t *testing.T) {
	arr := []int64{-4, 3, -9, 0, 4, 1}
	expected := []string{"0.500000", "0.333333", "0.166667"}

	result := plusMinus(arr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Plus minus first case was incorrect, got: %s, want: %s.", result, expected)
	}
}
