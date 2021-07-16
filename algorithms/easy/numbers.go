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

import (
	"sort"
)

// pickingNumbers return the maximum number of integers you can select from the array that the absolute difference
// between any two of the chosen integers is less than or equal to 1.
func pickingNumbers(a []int32) int32 {

	m := make(map[int32]int32)

	// Memorize array in order and count the number of each int.
	for i := range a {
		current := a[i]
		if val, ok := m[current]; ok {
			m[current] = val + 1
		} else {
			m[current] = 1
		}
	}

	// Basic case when there is one element.
	if len(m) == 1 {
		return int32(len(a))
	}

	// Sort the keys of the map to calculate maximum length.
	keys := make([]int32, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	maximumLength := m[keys[0]]

	// Iterate over map looking for integers which absolute difference is less than or equal to 1.
	for i := 0; i < (len(keys) - 1); i++ {
		if keys[i+1]-keys[i] <= 1 {
			currentLength := m[keys[i+1]] + m[keys[i]]

			if currentLength > maximumLength {
				maximumLength = currentLength
			}
		} else {
			currentLength := m[keys[i+1]]

			if currentLength > maximumLength {
				maximumLength = currentLength
			}
		}
	}

	return maximumLength
}

// main function provided by hacker rank to execute the code on platform. To run locally uncomment the function.
//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
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
//	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
//
//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var a []int32
//
//	for i := 0; i < int(n); i++ {
//		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
//		checkError(err)
//		aItem := int32(aItemTemp)
//		a = append(a, aItem)
//	}
//
//	result := pickingNumbers(a)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
