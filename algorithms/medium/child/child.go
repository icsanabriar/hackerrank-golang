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

// max retrieves the maximum value of the given values.
func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

// commonChild find the child string of the given ones, by eliminating zero or more characters.
func commonChild(s1 string, s2 string) int32 {

	lengthS1 := len(s1)
	lengthS2 := len(s2)

	var common [][]int32

	for i := 0; i <= lengthS1; i++ {
		var row []int32
		for j := 0; j <= lengthS2; j++ {
			row = append(row, 0)
		}
		common = append(common, row)
	}

	for i := 1; i <= lengthS1; i++ {
		for j := 1; j <= lengthS2; j++ {
			if s1[i-1] == s2[j-1] {
				common[i][j] = common[i-1][j-1] + 1
			} else {
				common[i][j] = max(common[i-1][j], common[i][j-1])
			}
		}
	}

	return common[lengthS1][lengthS2]
}
