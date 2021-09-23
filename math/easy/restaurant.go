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

// gcd calculate the great common divisor of given a,b.
func gcd(a int32, b int32) int32 {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// restaurant calculates the number of squares that could be build using the given l and b.
func restaurant(l int32, b int32) int32 {
	g := gcd(l, b)
	return (l * b) / (g * g)
}

// main function provided by hacker rank to execute the code on platform.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 1024*1024)
//
//	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	t := int32(tTemp)
//
//	for tItr := 0; tItr < int(t); tItr++ {
//		lb := strings.Split(readLine(reader), " ")
//
//		lTemp, err := strconv.ParseInt(lb[0], 10, 64)
//		checkError(err)
//		l := int32(lTemp)
//
//		bTemp, err := strconv.ParseInt(lb[1], 10, 64)
//		checkError(err)
//		b := int32(bTemp)
//
//		result := restaurant(l, b)
//
//		fmt.Fprintf(writer, "%d\n", result)
//	}
//
//	writer.Flush()
//}
