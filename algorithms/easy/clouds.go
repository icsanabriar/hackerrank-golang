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

// jumpingOnClouds return an integer representing the minimum number of jumps required to get to the end of given
// array c taking into account save jumps.
func jumpingOnClouds(c []int32) int32 {

	i := 0
	jumps := int32(0)

	for i < len(c) {
		if i+2 < len(c) {
			twoSteps := c[i+2]
			if twoSteps == 0 {
				i = i + 2
			} else {
				i = i + 1
			}
		} else {
			i = i + 1
		}
		jumps++
	}

	return jumps - 1
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
//
//	n := int32(nTemp)
//	cTemp := strings.Split(readLine(reader), " ")
//
//	var c []int32
//
//	for i := 0; i < int(n); i++ {
//		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
//		checkError(err)
//
//		cItem := int32(cItemTemp)
//		c = append(c, cItem)
//	}
//
//	result := jumpingOnClouds(c)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
