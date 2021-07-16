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

// findDigits return an integer representing the number of digits of d that are divisor of given n.
func findDigits(n int32) int {

	divisorDigit := false
	counter := 0
	originalN := n
	previousDigit := make(map[int32]bool)

	for n > 0 {
		currentDigit := n % 10

		if currentDigit != 0 {
			if val, ok := previousDigit[currentDigit]; !ok {

				divisorDigit = (originalN % currentDigit) == 0
				previousDigit[currentDigit] = divisorDigit

				if divisorDigit {
					counter += 1
				}

			} else if val {
				counter += 1
			}
		}

		n = n / 10
	}

	return counter
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
//	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	t := int32(tTemp)
//
//	for tItr := 0; tItr < int(t); tItr++ {
//		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		n := int32(nTemp)
//
//		result := findDigits(n)
//
//		_, _ = fmt.Fprintf(writer, "%d\n", result)
//	}
//
//	_ = writer.Flush()
//}
