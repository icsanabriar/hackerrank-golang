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

// hourglassSum calculate the maximum hourglass sum in the given arr.
func hourglassSum(arr [][]int32) int32 {

	maximum := int32(-64)
	var i, j int
	var sum int32

	// Iterate over rows of arr.
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {

			// First row of the glass.
			sum = arr[i][j]
			sum += arr[i][j+1]
			sum += arr[i][j+2]

			// Second row of the glass.
			sum += arr[i+1][j+1]

			// Third row of the glass.
			sum += arr[i+2][j]
			sum += arr[i+2][j+1]
			sum += arr[i+2][j+2]

			if maximum < sum {
				maximum = sum
			}
		}
	}

	return maximum
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
//	var arr [][]int32
//
//	for i := 0; i < 6; i++ {
//		arrRowTemp := strings.Split(readLine(reader), " ")
//		var arrRow []int32
//
//		for _, arrRowItem := range arrRowTemp {
//			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
//			checkError(err)
//			arrItem := int32(arrItemTemp)
//			arrRow = append(arrRow, arrItem)
//		}
//
//		if len(arrRow) != 6 {
//			panic("Bad input")
//		}
//
//		arr = append(arr, arrRow)
//	}
//
//	result := hourglassSum(arr)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
