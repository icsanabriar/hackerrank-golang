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

// permutationEquation return an array with the values of p(p(y)) base on the values stored given on p.
func permutationEquation(p []int32) []int32 {

	result := make([]int32, len(p))

	for i := range p {
		for j := range result {
			if p[p[j]-1] == int32(i+1) {
				result[i] = int32(j + 1)
				break
			}
		}
	}

	return result
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
//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var p []int32
//
//	for i := 0; i < int(n); i++ {
//		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
//		checkError(err)
//		pItem := int32(pItemTemp)
//		p = append(p, pItem)
//	}
//
//	result := permutationEquation(p)
//
//	for i, resultItem := range result {
//		_, _ = fmt.Fprintf(writer, "%d", resultItem)
//
//		if i != len(result) - 1 {
//			_, _ = fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
