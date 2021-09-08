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

// Sum the values of the given array.
func sumHeights(h []int32) int64 {
	total := int64(0)

	for i := 0; i < len(h); i++ {
		total = total + int64(h[i])
	}

	return total
}

// Poll element from array.
func pollElement(h1 *[]int32, h2 *[]int32, h3 *[]int32, index int) int64 {
	element := int32(0)

	if index == 0 && len(*h1) > 0 {
		element = (*h1)[0]
		*h1 = (*h1)[1:]
	} else if index == 1 && len(*h2) > 0 {
		element = (*h2)[0]
		*h2 = (*h2)[1:]
	} else if index == 2 && len(*h3) > 0 {
		element = (*h3)[0]
		*h3 = (*h3)[1:]
	}

	return int64(element)
}

// equalStacks calculate the maximum height the three given stacks could be equal.
func equalStacks(h1 []int32, h2 []int32, h3 []int32) int64 {
	cache := make([]int64, 0)

	cache = append(cache, sumHeights(h1))
	cache = append(cache, sumHeights(h2))
	cache = append(cache, sumHeights(h3))

	for !(cache[0] == cache[1] && cache[1] == cache[2]) {
		maxHeightIndex := 0
		for i := 1; i < len(cache); i++ {
			if cache[i] > cache[maxHeightIndex] {
				maxHeightIndex = i
			}
		}
		cache[maxHeightIndex] = cache[maxHeightIndex] - pollElement(&h1, &h2, &h3, maxHeightIndex)
	}

	return cache[0]
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
//	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	n1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//	checkError(err)
//	n1 := int32(n1Temp)
//
//	n2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	n2 := int32(n2Temp)
//
//	n3Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
//	checkError(err)
//	n3 := int32(n3Temp)
//
//	h1Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var h1 []int32
//
//	for i := 0; i < int(n1); i++ {
//		h1ItemTemp, err := strconv.ParseInt(h1Temp[i], 10, 64)
//		checkError(err)
//		h1Item := int32(h1ItemTemp)
//		h1 = append(h1, h1Item)
//	}
//
//	h2Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var h2 []int32
//
//	for i := 0; i < int(n2); i++ {
//		h2ItemTemp, err := strconv.ParseInt(h2Temp[i], 10, 64)
//		checkError(err)
//		h2Item := int32(h2ItemTemp)
//		h2 = append(h2, h2Item)
//	}
//
//	h3Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var h3 []int32
//
//	for i := 0; i < int(n3); i++ {
//		h3ItemTemp, err := strconv.ParseInt(h3Temp[i], 10, 64)
//		checkError(err)
//		h3Item := int32(h3ItemTemp)
//		h3 = append(h3, h3Item)
//	}
//
//	result := equalStacks(h1, h2, h3)
//
//	_, _ = fmt.Fprintf(writer, "%d\n", result)
//
//	_ = writer.Flush()
//}
