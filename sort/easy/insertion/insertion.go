// Copyright 2022 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"io"
)

// insertionSort print sorted values of the given array.
func insertionSort(w io.Writer, n int, arr []int64) {
	temp := arr[n-1]
	index := n - 2

	for index >= 0 && arr[index] > temp {
		arr[index+1] = arr[index]
		printArray(w, arr)
		index--
	}

	arr[index+1] = temp
	printArray(w, arr)
}

// printArray prints the given array into the console.
func printArray(w io.Writer, arr []int64) {
	for _, v := range arr {
		_, _ = fmt.Fprintf(w, "%d ", v)
	}
	_, _ = fmt.Fprintln(w, "")
}
