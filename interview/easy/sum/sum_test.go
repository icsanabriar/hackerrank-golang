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
	"testing"
)

// TestMinMaxSumFirstGivenCase implements the test given as first example on hackerrank.
func TestMinMaxSumFirstGivenCase(t *testing.T) {

	arr := []int32{1, 3, 5, 7, 9}
	emin := int64(16)
	emax := int64(24)

	rmin, rmax := miniMaxSum(arr)

	if rmin != emin || rmax != emax {
		t.Errorf("Min Max Sum first case was incorrect, got: %d, want: %d.", rmin, emin)
		t.Errorf("Min Max Sum first case was incorrect, got: %d, want: %d.", rmax, emax)
	}
}

// TestMinMaxSumSecondGivenCase implements the test given as second example on hackerrank.
func TestMinMaxSumSecondGivenCase(t *testing.T) {

	arr := []int32{7, 69, 2, 221, 8974}
	emin := int64(299)
	emax := int64(9271)

	rmin, rmax := miniMaxSum(arr)

	if rmin != emin || rmax != emax {
		t.Errorf("Min Max Sum second case was incorrect, got: %d, want: %d.", rmin, emin)
		t.Errorf("Min Max Sum second case was incorrect, got: %d, want: %d.", rmax, emax)
	}
}
