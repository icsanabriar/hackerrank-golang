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

// TestTimeConversionFirstGivenCase implements the test given as first example on hackerrank.
func TestTimeConversionFirstGivenCase(t *testing.T) {

	s := "12:01:00PM"
	expected := "12:01:00"

	result := timeConversion(s)

	if result != expected {
		t.Errorf("Time conversion first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestTimeConversionSecondGivenCase implements the test given as second example on hackerrank.
func TestTimeConversionSecondGivenCase(t *testing.T) {

	s := "12:01:00AM"
	expected := "00:01:00"

	result := timeConversion(s)

	if result != expected {
		t.Errorf("Time conversion second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestTimeConversionThirdGivenCase implements the test given as third example on hackerrank.
func TestTimeConversionThirdGivenCase(t *testing.T) {

	s := "07:05:45PM"
	expected := "19:05:45"

	result := timeConversion(s)

	if result != expected {
		t.Errorf("Time conversion third case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestTimeConversionFourthGivenCase implements the test given as fourth example on hackerrank.
func TestTimeConversionFourthGivenCase(t *testing.T) {

	s := "02:07:33AM"
	expected := "02:07:33"

	result := timeConversion(s)

	if result != expected {
		t.Errorf("Time conversion fourth case was incorrect, got: %s, want: %s.", result, expected)
	}
}
