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

// hurdleRace return an integer representing the minimum number of doses required to jump all of the hurdle.
func hurdleRace(k int32, height []int32) int32 {

	max := height[0]

	for _, value := range height {
		if value > max {
			max = value
		}
	}

	if (max - k) < 0 {
		return 0
	}

	return max - k
}

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
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
//	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	k := int32(kTemp)
//
//	heightTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var height []int32
//
//	for i := 0; i < int(n); i++ {
//		heightItemTemp, err := strconv.ParseInt(heightTemp[i], 10, 64)
//		checkError(err)
//
//		heightItem := int32(heightItemTemp)
//		height = append(height, heightItem)
//	}
//
//	result := hurdleRace(k, height)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
