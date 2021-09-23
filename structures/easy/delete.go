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

// SinglyLinkedListNode provided on hackerrank.
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// SinglyLinkedList provided on hackerrank.
type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

// insertNodeIntoSinglyLinkedList provided on hackerrank.
func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
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

// deleteNode returns the node is being deleted from the given structure.
func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {

	tempHead := SinglyLinkedListNode{
		next: llist,
	}

	temp := &tempHead

	for i := 0; i < int(position); i++ {
		temp = temp.next
	}

	temp.next = temp.next.next

	return tempHead.next
}

// printSinglyLinkedList provided on hackerrank.
//func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
//	for node != nil {
//		_, _ = fmt.Fprintf(writer, "%d", node.data)
//
//		node = node.next
//
//		if node != nil {
//			_, _ = fmt.Fprintf(writer, sep)
//		}
//	}
//}
//
// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer func(stdout *os.File) {
//		err := stdout.Close()
//		if err != nil {
//			fmt.Println("Error reading output path file.")
//		}
//	}(stdout)
//
//	writer := bufio.NewWriterSize(stdout, 1024 * 1024)
//
//	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	llist := SinglyLinkedList{}
//	for i := 0; i < int(llistCount); i++ {
//		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		llistItem := int32(llistItemTemp)
//		llist.insertNodeIntoSinglyLinkedList(llistItem)
//	}
//
//	positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	position := int32(positionTemp)
//
//	llist1 := deleteNode(llist.head, position)
//
//	printSinglyLinkedList(llist1, " ", writer)
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
