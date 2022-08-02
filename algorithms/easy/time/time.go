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
	"strconv"
	"strings"
)

// timeConversion converts the given string that contains time in 12-hour AM/PM format to military 24-hour time.
func timeConversion(s string) string {

	// Split by Colon.
	colon := ":"
	split := strings.Split(s, colon)

	// Define AM / PM.
	hour, _ := strconv.Atoi(split[0])

	if strings.HasSuffix(split[2], "AM") {

		// Update Midnight hour.
		if hour == 12 {
			split[0] = "00"
		}

	} else {

		// Add 12 hours to time.
		if hour != 12 {
			hour = hour + 12
		}

		split[0] = fmt.Sprintf("%02d", hour)
	}

	return split[0] + colon + split[1] + colon + split[2][0:2]
}
