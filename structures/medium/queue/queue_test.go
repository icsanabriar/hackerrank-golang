// Copyright 2020 Iván Camilo Sanabria
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

// TestQueueGivenCase implements the test given as first example on hackerrank.
func TestQueueGivenCase(t *testing.T) {

	var q = queue{}

	var input = int32(42)
	q.queueElement(input)

	if input != q.firstElement() {
		t.Errorf("Queue element first input was incorrect, got: %v, want: %v.", q.firstElement(), input)
	}

	q.dequeueElement()

	if 0 != q.length() {
		t.Errorf("Dequeue element second input was incorrect, got: %v, want: %v.", q.length(), 0)
	}

	input = int32(14)
	q.queueElement(input)

	if input != q.firstElement() {
		t.Errorf("Queue element third input was incorrect, got: %v, want: %v.", q.firstElement(), input)
	}

	first := q.firstElement()
	expected := int32(14)

	if expected != first {
		t.Errorf("First element fourth input was incorrect, got: %v, want: %v.", first, expected)
	}

	input = int32(28)
	q.queueElement(input)

	if 2 != q.length() {
		t.Errorf("Queue element fifth input was incorrect, got: %v, want: %v.", q.length(), 2)
	}

	first = q.firstElement()
	expected = int32(14)

	if expected != first {
		t.Errorf("First element sixth input was incorrect, got: %v, want: %v.", first, expected)
	}

	input = int32(60)
	q.queueElement(input)

	if 3 != q.length() {
		t.Errorf("Queue element seventh input was incorrect, got: %v, want: %v.", q.length(), 2)
	}

	input = int32(78)
	q.queueElement(input)

	if 4 != q.length() {
		t.Errorf("Queue element eigth input was incorrect, got: %v, want: %v.", q.length(), 2)
	}

	q.dequeueElement()

	if 3 != q.length() {
		t.Errorf("Dequeue element ninth input was incorrect, got: %v, want: %v.", q.length(), 3)
	}

	q.dequeueElement()

	if 2 != q.length() {
		t.Errorf("Dequeue element ninth input was incorrect, got: %v, want: %v.", q.length(), 2)
	}
}
