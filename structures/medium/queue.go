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

// queue structure compose by two stacks.
type queue struct {
	recipient []int32
	messenger []int32
}

// queueElement is responsible of queuing the given x into recipient stack.
func (q *queue) queueElement(x int32) {
	q.recipient = append(q.recipient, x)
}

// transfer is responsible of moving elements from recipient to messenger stack.
func (q *queue) transferElements() {
	if len(q.messenger) == 0 {
		for len(q.recipient) > 0 {
			q.messenger = append(q.messenger, q.recipient[len(q.recipient)-1])
			q.recipient = q.recipient[:len(q.recipient)-1]
		}
	}
}

// dequeueElement is responsible of removing the first element of the queue using transfer.
func (q *queue) dequeueElement() {
	q.transferElements()
	q.messenger = q.messenger[:len(q.messenger)-1]
}

// firstElement is responsible of getting the first element of the queue.
func (q *queue) firstElement() int32 {
	q.transferElements()
	return q.messenger[len(q.messenger)-1]
}

// length is responsible for getting the size of the stack.
func (q *queue) length() int {
	return len(q.recipient) + len(q.messenger)
}

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
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
//	writer := bufio.NewWriterSize(stdout, 1024*1024)
//
//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	var q = queue{}
//
//	for i := 0; i < int(n); i++ {
//		command := strings.Split(readLine(reader), " ")
//
//		// Command 1 is the only one with parameters.
//		if len(command) == 2 {
//			xTemp, err := strconv.ParseInt(command[1], 10, 64)
//			checkError(err)
//			x := int32(xTemp)
//			// Queue element into recipient stack.
//			q.queueElement(x)
//		} else {
//			// Check command 2 to dequeue element.
//			if command[0] == "2" {
//				q.dequeueElement()
//			} else {
//				// Print the first element on the queue.
//				_, _ = fmt.Fprintf(writer, "%d\n", q.firstElement())
//			}
//		}
//	}
//
//	_ = writer.Flush()
//}
