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

import (
	"testing"
)

// TestDeleteNodeGivenCase implements the test given as first example on hackerrank.
func TestDeleteNodeGivenCase(t *testing.T) {
	n := 3
	array := []int64{20, 6, 2, 19, 7, 4, 15, 9}
	list := buildList(array)

	expected := int64(19)
	result := deleteNode(list.head, n)

	for result != nil {
		if expected == result.data {
			t.Errorf("Delete node first case was incorrect, got: %d expected to be removed.", result.data)
		}

		result = result.next
	}
}

// buildList factory to create a SinglyLinkedList from a given array.
func buildList(array []int64) SinglyLinkedList {
	list := SinglyLinkedList{}
	for i := 0; i < len(array); i++ {
		list.insertNodeIntoSinglyLinkedList(array[i])
	}

	return list
}
