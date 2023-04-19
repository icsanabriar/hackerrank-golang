// Copyright 2023 Iván Camilo Sanabria
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
	"strconv"
	"strings"
)

// timeConversion format the given time into 24 hours.
func timeConversion(s string) string {

	hour, _ := strconv.Atoi(s[0:2])

	if strings.HasSuffix(s, "PM") && hour != 12 {
		hour += 12
	} else if strings.HasSuffix(s, "AM") && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%s", hour, s[3:len(s)-2])
}
