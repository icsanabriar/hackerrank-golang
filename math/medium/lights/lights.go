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

// modulus is the value to simplify the result.
const modulus = int32(100000)

// lights calculate the total number of patterns that could be made using the given N serial light bulbs modulo 10e5.
func lights(n int32) int32 {
	total := int32(1)

	for i := int32(0); i < n; i++ {
		total = 2 * total % modulus
	}

	// Total - 1 to complete the series, then add the modulus to get original value and apply module to obtain result.
	return (total - 1 + modulus) % modulus
}
