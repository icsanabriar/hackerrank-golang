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

// largestRectangle return an integer representing the largest rectangle that can be formed within the bounds of
// consecutive buildings given the heights on array h.
func largestRectangle(h []int32) int64 {

	largest := int64(0)
	var nearest int32

	for i, hi := range h {
		nearest = 1

		// Go to the right and check nearest.
		for j := i + 1; j < len(h); j++ {
			if hi > h[j] {
				break
			} else {
				nearest += 1
			}
		}

		// Go to the left and check nearest.
		for j := i - 1; j >= 0; j-- {
			if hi > h[j] {
				break
			} else {
				nearest += 1
			}
		}

		area := int64(hi * nearest)

		if area > largest {
			largest = area
		}
	}

	return largest
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
//	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	hTemp := strings.Split(readLine(reader), " ")
//
//	var h []int32
//
//	for i := 0; i < int(n); i++ {
//		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
//		checkError(err)
//		hItem := int32(hItemTemp)
//		h = append(h, hItem)
//	}
//
//	result := largestRectangle(h)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
