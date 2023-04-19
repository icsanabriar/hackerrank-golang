// Copyright 2020 Iván Camilo Sanabria
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

// countingValleys calculates how many valleys were walked by Gary on the given road s.
func countingValleys(s string) int32 {

	counter := int32(0)
	level := int32(0)

	for _, c := range s {
		previousLevel := level

		if string(c) == "U" {
			level++
		} else {
			level--
		}

		if previousLevel < 0 && level == 0 {
			counter++
		}
	}

	return counter
}
