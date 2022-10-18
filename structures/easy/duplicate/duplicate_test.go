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

// TestGetNodeValueGivenCase implements the test given as first example on hackerrank.
func TestGetNodeValueGivenCase(t *testing.T) {

	list := SinglyLinkedList{}
	list.insertNodeIntoSinglyLinkedList(int32(3))
	list.insertNodeIntoSinglyLinkedList(int32(3))
	list.insertNodeIntoSinglyLinkedList(int32(3))
	list.insertNodeIntoSinglyLinkedList(int32(4))
	list.insertNodeIntoSinglyLinkedList(int32(5))
	list.insertNodeIntoSinglyLinkedList(int32(5))

	expected := &SinglyLinkedList{}
	expected.insertNodeIntoSinglyLinkedList(int32(3))
	expected.insertNodeIntoSinglyLinkedList(int32(4))
	expected.insertNodeIntoSinglyLinkedList(int32(5))
	node := expected.head

	result := removeDuplicates(list.head)

	for node.next != nil {
		if result.data != node.data {
			t.Errorf("Get node value first case was incorrect, got: %d, want: %d.", result.data, node.data)
		}

		result = result.next
		node = node.next
	}
}
