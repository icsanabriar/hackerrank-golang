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

import "strconv"

// maxBribes defines maximum constant of maximum bribes allowed.
const maxBribes = int32(2)

// max returns the largest of a or b.
func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

// minimumBribes calculate the number of bribes necessary to get q in the order that is given as parameter.
func minimumBribes(q []int32) string {

	numberBribes := int32(0)

	for i := len(q) - 1; i >= 0; i-- {
		marker := int32(i + 1)

		if q[i]-marker > maxBribes {
			return "Too chaotic"
		}

		lowerBound := max(q[i]-2, 0)
		upperBound := max(int32(i-1), 0)

		for j := upperBound; j >= lowerBound; j-- {
			if q[j] > q[i] {
				numberBribes++
			}
		}
	}

	return strconv.FormatInt(int64(numberBribes), 10)
}
