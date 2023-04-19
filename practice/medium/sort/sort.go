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

// generateRange retrieves the range to count first indexes.
func generateRange(min int32, max int32, expenditure []int32, d int32) []int32 {
	a := make([]int32, max-min+1)

	for i := int32(0); i < d; i++ {
		a[expenditure[i]]++
	}

	return a
}

// activityNotifications validates that the expenditure is not more than double of the median given series of size d.
func activityNotifications(expenditure []int32, d int32) int32 {
	counter := int32(0)
	counts := generateRange(0, 200, expenditure, d)

	for i := d; i < int32(len(expenditure)); i++ {
		lower := int32(0)
		left := int32(0)

		for ((left + counts[lower]) * 2) <= d {
			left = left + counts[lower]
			lower++
		}

		upper := int32(len(counts) - 1)
		right := int32(0)

		for ((right + counts[upper]) * 2) <= d {
			right = right + counts[upper]
			upper--
		}

		if expenditure[i] >= (lower + upper) {
			counter++
		}

		counts[expenditure[i-d]]--
		counts[expenditure[i]]++
	}

	return counter
}
