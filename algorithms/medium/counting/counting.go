// Copyright 2022 Iván Camilo Sanabria
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
	"fmt"
	"io"
)

//countSort prints the given array in order taking into account the given index and value on the array.
func countSort(w io.Writer, input map[int64]string) {

	cache := make([]int64, 100)

	for i := range cache {
		cache[i] = int64(i)
	}

	for _, key := range cache {
		if value, ok := input[key]; ok {
			_, _ = fmt.Fprintf(w, "%s ", value)
		}
	}

	_, _ = fmt.Fprintln(w, "")
}
