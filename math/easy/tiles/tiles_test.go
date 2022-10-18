// Copyright 2021 Iván Camilo Sanabria
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

import (
	"reflect"
	"testing"
)

// TestSherlockMovingTilesGivenCase implements the test given as first example on hackerrank.
func TestSherlockMovingTilesGivenCase(t *testing.T) {

	l := int32(10)
	s1 := int32(1)
	s2 := int32(2)
	queries := []int64{50, 100}

	expected := []float64{4.142135623730949, 0.0000}

	result := movingTiles(l, s1, s2, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sherlock moving tiles first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestSherlockMovingTilesHiddenGivenCase implements the test given as hidden example on hackerrank.
func TestSherlockMovingTilesHiddenGivenCase(t *testing.T) {

	l := int32(10000000)
	s1 := int32(1)
	s2 := int32(99999997)
	queries := []int64{81258385599125, 16248989464296, 47115376367013}

	expected := []float64{0.01393929091961771879, 0.08441436098337191940, 0.044348832168970016}

	result := movingTiles(l, s1, s2, queries)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sherlock moving tiles hidden case was incorrect, got: %v, want: %v.", result, expected)
	}
}
