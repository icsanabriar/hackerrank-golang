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

// permutationEquation return an array with the values of p(p(y)) base on the values stored given on p.
func permutationEquation(p []int32) []int32 {

	result := make([]int32, len(p))

	for i := range p {
		for j := range result {
			if p[p[j]-1] == int32(i+1) {
				result[i] = int32(j + 1)
				break
			}
		}
	}

	return result
}
