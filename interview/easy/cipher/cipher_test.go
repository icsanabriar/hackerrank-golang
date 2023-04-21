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

// TestCaesarCipherFirstGivenCase implements the test given as first example on hackerrank.
func TestCaesarCipherFirstGivenCase(t *testing.T) {

	s := "middle-Outz"
	k := int32(2)
	expected := "okffng-Qwvb"

	result := caesarCipher(s, k)

	if result != expected {
		t.Errorf("Caesar chiper first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestCaesarCipherSecondGivenCase implements the test given as second example on hackerrank.
func TestCaesarCipherSecondGivenCase(t *testing.T) {

	s := "Always-Look-on-the-Bright-Side-of-Life"
	k := int32(5)
	expected := "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj"

	result := caesarCipher(s, k)

	if result != expected {
		t.Errorf("Caesar chiper second case was incorrect, got: %s, want: %s.", result, expected)
	}
}
