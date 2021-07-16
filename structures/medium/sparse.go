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

// matchingStrings count the given queries on the string arrays. The query and string should be the same.
func matchingStrings(strings []string, queries []string) []int32 {

	cache := make(map[string]int32)

	// Count the number of strings and store in a map.
	for _, e := range strings {

		if val, ok := cache[e]; ok {
			cache[e] = val + 1
		} else {
			cache[e] = 1
		}
	}

	result := make([]int32, 0)

	// Find matches of queries.
	for _, e := range queries {
		result = append(result, cache[e])
	}

	return result
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
//	stringsCount, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	var strings []string
//
//	for i := 0; i < int(stringsCount); i++ {
//		stringsItem := readLine(reader)
//		strings = append(strings, stringsItem)
//	}
//
//	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	var queries []string
//
//	for i := 0; i < int(queriesCount); i++ {
//		queriesItem := readLine(reader)
//		queries = append(queries, queriesItem)
//	}
//
//	res := matchingStrings(strings, queries)
//
//	for i, resItem := range res {
//		_, _ = fmt.Fprintf(writer, "%d", resItem)
//
//		if i != len(res)-1 {
//			_, _ = fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	_, _ = fmt.Fprintf(writer, "\n")
//
//	_ = writer.Flush()
//}
