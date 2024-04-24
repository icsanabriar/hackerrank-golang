// Copyright 2022 Iván Camilo Sanabria
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

// TestGetNodeValueFirstGivenCase implements the test given as first example on hackerrank.
func TestGetNodeValueFirstGivenCase(t *testing.T) {
	position := int64(0)
	list := SinglyLinkedList{}
	list.insertNodeIntoSinglyLinkedList(int64(1))
	expected := int64(1)

	result := getNode(list.head, position)

	if result != expected {
		t.Errorf("Get node value first case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestGetNodeValueSecondGivenCase implements the test given as second example on hackerrank.
func TestGetNodeValueSecondGivenCase(t *testing.T) {
	position := int64(2)
	list := SinglyLinkedList{}
	list.insertNodeIntoSinglyLinkedList(int64(3))
	list.insertNodeIntoSinglyLinkedList(int64(2))
	list.insertNodeIntoSinglyLinkedList(int64(1))
	expected := int64(3)

	result := getNode(list.head, position)

	if result != expected {
		t.Errorf("Get node value second case was incorrect, got: %d, want: %d.", result, expected)
	}
}

// TestGetNodeValueThirdGivenCase implements the test given as third example on hackerrank.
func TestGetNodeValueThirdGivenCase(t *testing.T) {
	position := int64(1)
	list := SinglyLinkedList{}
	list.insertNodeIntoSinglyLinkedList(int64(3))
	list.insertNodeIntoSinglyLinkedList(int64(2))
	list.insertNodeIntoSinglyLinkedList(int64(1))
	list.insertNodeIntoSinglyLinkedList(int64(5))
	list.insertNodeIntoSinglyLinkedList(int64(6))
	list.insertNodeIntoSinglyLinkedList(int64(3))
	expected := int64(6)

	result := getNode(list.head, position)

	if result != expected {
		t.Errorf("Get node value third case was incorrect, got: %d, want: %d.", result, expected)
	}
}
