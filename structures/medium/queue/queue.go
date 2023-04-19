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

// queue structure compose by two stacks.
type queue struct {
	recipient []int32
	messenger []int32
}

// queueElement is responsible of queuing the given x into recipient stack.
func (q *queue) queueElement(x int32) {
	q.recipient = append(q.recipient, x)
}

// transfer is responsible for moving elements from recipient to messenger stack.
func (q *queue) transferElements() {
	if len(q.messenger) == 0 {
		for len(q.recipient) > 0 {
			q.messenger = append(q.messenger, q.recipient[len(q.recipient)-1])
			q.recipient = q.recipient[:len(q.recipient)-1]
		}
	}
}

// dequeueElement is responsible for removing the first element of the queue using transfer.
func (q *queue) dequeueElement() {
	q.transferElements()
	q.messenger = q.messenger[:len(q.messenger)-1]
}

// firstElement is responsible for getting the first element of the queue.
func (q *queue) firstElement() int32 {
	q.transferElements()
	return q.messenger[len(q.messenger)-1]
}

// length is responsible for getting the size of the stack.
func (q *queue) length() int {
	return len(q.recipient) + len(q.messenger)
}
