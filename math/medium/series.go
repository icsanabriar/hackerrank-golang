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

// module is the value to simplify the result.
const module = 1000000007

// summingSeries calculates the sum of specific series when the nth term is n2 - (n-1)2.
// Simplifying the expression it is the sum of nth odd numbers which is equal a n2.
func summingSeries(n int64) int64 {

	n %= module

	return n * n % module
}

// main function provided by hacker rank to execute the code on platform.
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
//		n, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//
//		result := summingSeries(n)
//
//		_, _ = fmt.Fprintf(writer, "%d\n", result)
//	}
//
//	_ = writer.Flush()
//}
