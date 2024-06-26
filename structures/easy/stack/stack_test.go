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

// TestStackGivenCase implements the test given as first example on hackerrank.
func TestStackGivenCase(t *testing.T) {
	var s = stack{}
	var maxStack = stack{}

	var input = int64(91)
	s.push(input)
	maxStack.pushMax(input)

	if input != s.peek() {
		t.Errorf("Stack push first input was incorrect, got: %v, want: %v.", s.peek(), input)
	}

	s.pop()
	maxStack.pop()

	if s.Length() != 0 {
		t.Errorf("Stack pop second input was incorrect, got: %v, want: %v.", s.Length(), 0)
	}

	input = int64(20)
	s.push(input)
	maxStack.pushMax(input)

	if input != s.peek() {
		t.Errorf("Stack push third input was incorrect, got: %v, want: %v.", s.peek(), input)
	}

	s.pop()
	maxStack.pop()

	if s.Length() != 0 {
		t.Errorf("Stack pop fourth input was incorrect, got: %v, want: %v.", s.Length(), 0)
	}

	input = int64(26)
	s.push(input)
	maxStack.pushMax(input)

	if input != s.peek() {
		t.Errorf("Stack push fifth input was incorrect, got: %v, want: %v.", s.peek(), input)
	}

	input = int64(20)
	s.push(input)
	maxStack.pushMax(input)

	if input != s.peek() {
		t.Errorf("Stack push sixth input was incorrect, got: %v, want: %v.", s.peek(), input)
	}

	s.pop()
	maxStack.pop()

	if s.Length() != 1 {
		t.Errorf("Stack pop seventh input was incorrect, got: %v, want: %v.", s.Length(), 1)
	}

	maximum := maxStack.peek()
	expected := int64(26)

	if expected != maximum {
		t.Errorf("Stack print maximum element eight input was incorrect, got: %v, want: %v.", maximum, expected)
	}

	input = int64(91)
	s.push(input)
	maxStack.pushMax(input)

	if input != s.peek() {
		t.Errorf("Stack push ninth input was incorrect, got: %v, want: %v.", s.peek(), input)
	}

	maximum = maxStack.peek()
	expected = int64(91)

	if expected != maximum {
		t.Errorf("Stack print maximum element tenth input was incorrect, got: %v, want: %v.", maximum, expected)
	}
}
