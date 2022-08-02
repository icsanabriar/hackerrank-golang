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

// stack structure compose by an array.
type stack struct {
	arr []int32
}

// push is responsible for adding the element to the stack.
func (s *stack) push(e int32) {
	s.arr = append(s.arr, e)
}

// pushMax is responsible for adding the element to the stack.
func (s *stack) pushMax(e int32) {

	max := e

	if len(s.arr) > 0 && max < s.peek() {
		max = s.peek()
	}

	s.arr = append(s.arr, max)
}

// pop is responsible for deleting the top element of the stack.
func (s *stack) pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

// peek is responsible for getting the element at the top of the stack.
func (s *stack) peek() int32 {
	return s.arr[len(s.arr)-1]
}

// length is responsible for getting the size of the stack.
func (s *stack) length() int {
	return len(s.arr)
}
