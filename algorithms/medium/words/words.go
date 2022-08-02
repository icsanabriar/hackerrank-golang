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

// numberString defines static map for numbers in string.
var numberString = map[int32]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "quarter",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	21: "twenty one",
	22: "twenty two",
	23: "twenty three",
	24: "twenty four",
	25: "twenty five",
	26: "twenty six",
	27: "twenty seven",
	28: "twenty eight",
	29: "twenty nine",
	30: "half",
}

// timeInWords return a string with the given hour and minute using o'clock, past and to.
func timeInWords(h int32, m int32) string {

	minutesString := " minute"

	// If minutes is higher than 10 is plural.
	if m > 10 {
		minutesString = " minutes"
	}

	// Validate minutes should be include in time.
	if m%15 == 0 {
		minutesString = ""
	}

	// O'clock time output.
	if m == 0 {
		return numberString[h] + " o' clock"
	}

	// Past time output.
	if 0 < m && m < 31 {
		return numberString[m] + minutesString + " past " + numberString[h]
	}

	// To time output assuming the preconditions of m are met.
	hourTo := int32(1)

	if h < 12 {
		hourTo = h + 1
	}

	return numberString[60-m] + minutesString + " to " + numberString[hourTo]
}
