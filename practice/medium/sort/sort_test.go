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

// TestFraudulentActivityFirstGivenCase implements the test given as first example on hackerrank.
func TestFraudulentActivityFirstGivenCase(t *testing.T) {

	input := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	d := int32(5)
	expected := int32(2)

	result := activityNotifications(input, d)

	if result != expected {
		t.Errorf("Fraudulent activity notifications first case was incorrect, got: %v, want: %v.", result, expected)
	}
}

// TestFraudulentActivitySecondGivenCase implements the test given as second example on hackerrank.
func TestFraudulentActivitySecondGivenCase(t *testing.T) {

	input := []int32{1, 2, 3, 4, 4}
	d := int32(4)
	expected := int32(0)

	result := activityNotifications(input, d)

	if result != expected {
		t.Errorf("Fraudulent activity notifications second case was incorrect, got: %v, want: %v.", result, expected)
	}
}
