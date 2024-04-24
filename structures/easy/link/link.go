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

// SinglyLinkedListNode structure to simulate linked list node.
type SinglyLinkedListNode struct {
	data int64
	next *SinglyLinkedListNode
}

// SinglyLinkedList structure to simulate linked list.
type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

// insertNodeIntoSinglyLinkedList insert a node into the list.
func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int64) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

// getNode retrieves the data using the head and given position from tail.
func getNode(list *SinglyLinkedListNode, positionFromTail int64) int64 {
	temp := list
	maxIndex := int64(0)

	for temp.next != nil {
		maxIndex++
		temp = temp.next
	}

	temp = list

	for i := int64(0); i < (maxIndex - positionFromTail); i++ {
		temp = temp.next
	}

	return temp.data
}
