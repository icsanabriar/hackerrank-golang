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
func insertionSort(w io.Writer, arr []int32) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = value
	}

	printArray(w, arr)
}

// printArray prints the given array into the console.
func printArray(w io.Writer, arr []int32) {
	for _, v := range arr {
		_, _ = fmt.Fprintf(w, "%d ", v)
	}
	_, _ = fmt.Fprintln(w, "")
}
