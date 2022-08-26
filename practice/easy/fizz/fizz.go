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
	"fmt"
	"io"
)

//fizzBuzz function that print series fizz buzz based on the number n.
func fizzBuzz(w io.Writer, n int32) {
	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			_, _ = fmt.Fprintln(w, "FizzBuzz")
		} else {
			if i%3 == 0 {
				_, _ = fmt.Fprintln(w, "Fizz")
			}
			if i%5 == 0 {
				_, _ = fmt.Fprintln(w, "Buzz")
			}
			if i%3 != 0 && i%5 != 0 {
				_, _ = fmt.Fprintln(w, i)
			}
		}
	}
}
