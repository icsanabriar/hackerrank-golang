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

import "testing"

// TestLeonardoPrimeGameFirstGivenCase implements the test given as first example on hackerrank.
func TestLeonardoPrimeGameFirstGivenCase(t *testing.T) {

	input := int64(1)
	expected := int32(0)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLeonardoPrimeGameSecondGivenCase implements the test given as second example on hackerrank.
func TestLeonardoPrimeGameSecondGivenCase(t *testing.T) {

	input := int64(2)
	expected := int32(1)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLeonardoPrimeGameThirdGivenCase implements the test given as third example on hackerrank.
func TestLeonardoPrimeGameThirdGivenCase(t *testing.T) {

	input := int64(3)
	expected := int32(1)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor third case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLeonardoPrimeGameFourthGivenCase implements the test given as fourth example on hackerrank.
func TestLeonardoPrimeGameFourthGivenCase(t *testing.T) {

	input := int64(500)
	expected := int32(4)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor fourth case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLeonardoPrimeGameFifthGivenCase implements the test given as fifth example on hackerrank.
func TestLeonardoPrimeGameFifthGivenCase(t *testing.T) {

	input := int64(5000)
	expected := int32(5)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor fifth case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestLeonardoPrimeGameSixthGivenCase implements the test given as sixth example on hackerrank.
func TestLeonardoPrimeGameSixthGivenCase(t *testing.T) {

	input := int64(10000000000)
	expected := int32(10)

	result := primeCount(input)

	if result != expected {
		t.Errorf("Leonardo's prime factor sixth case was incorrect, got: %d, want: %d.", result, expected)
	}
}
