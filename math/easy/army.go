// Copyright 2021 Iván Camilo Sanabria
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

// gameWithCells calculates how many packages should be dropped to supply all army.
func gameWithCells(n int32, m int32) int32 {
	return ((n + 1) / 2) * ((m + 1) / 2)
}

// main function provided by hacker rank to execute the code on platform.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
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
//	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
//
//	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	m := int32(mTemp)
//
//	result := gameWithCells(n, m)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
