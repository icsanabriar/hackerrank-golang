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

import "container/heap"

// Item define the object inside the priority queue.
type Item struct {
	priority int64
	index    int
}

// PriorityQueue structure used to store cookies array.
type PriorityQueue []*Item

// Len return the length of the priority queue.
func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// Less return a boolean based on the comparison between items on index i,j.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Pop return the item that is next based on expiration and remove from queue.
func (pq *PriorityQueue) Pop() interface{} {
	previous := *pq
	n := len(previous)

	item := previous[n-1]
	previous[n-1] = nil
	item.index = -1

	*pq = previous[:n-1]

	return item
}

// Push add the given item into the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)

	item := x.(*Item)
	item.index = n

	*pq = append(*pq, item)
}

// Swap exchange the elements using the given indexes.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// cookies estimates the number of operations required to create a special cookie.
func cookies(k int64, arr []int64) int64 {
	pq := make(PriorityQueue, len(arr))

	for i, item := range arr {
		pq[i] = &Item{
			priority: item,
			index:    i,
		}
	}

	heap.Init(&pq)

	operations := int64(0)
	lessSweet := heap.Pop(&pq).(*Item)

	for lessSweet.priority < k && pq.Len() >= 1 {
		secondLess := heap.Pop(&pq).(*Item)
		priority := lessSweet.priority + 2*(secondLess.priority)

		item := &Item{
			priority: priority,
		}

		heap.Push(&pq, item)

		lessSweet = heap.Pop(&pq).(*Item)
		operations++
	}

	if lessSweet.priority >= k {
		return operations
	}

	return -1
}
