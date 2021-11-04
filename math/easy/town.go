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

// connectingTowns calculates the total number of routes to connect the first to last town.
func connectingTowns(n int32, routes []int32) int32 {

	modulo := int32(1234567)
	total := int32(1)

	for i := int32(0); i < (n - 1); i++ {
		total = (total * routes[i]) % modulo
	}

	return total
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
//	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	t := int32(tTemp)
//
//	for tItr := 0; tItr < int(t); tItr++ {
//		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//		checkError(err)
//		n := int32(nTemp)
//
//		routesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//		var routes []int32
//
//		for i := 0; i < int(n) - 1; i++ {
//			routesItemTemp, err := strconv.ParseInt(routesTemp[i], 10, 64)
//			checkError(err)
//			routesItem := int32(routesItemTemp)
//			routes = append(routes, routesItem)
//		}
//
//		result := connectingTowns(n, routes)
//
//		_, _ = fmt.Fprintf(writer, "%d\n", result)
//	}
//
//	_ = writer.Flush()
//}
