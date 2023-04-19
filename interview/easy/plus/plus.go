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

import "fmt"

// plusMinus retrieves the ratios of positive, negative and zero values of the given array.
func plusMinus(arr []int32) []string {

	positive, negative, zeros := 0, 0, 0

	for _, value := range arr {
		switch {
		case value == 0:
			zeros++
		case value > 0:
			positive++
		default:
			negative++
		}
	}

	length := float32(len(arr))

	return []string{
		fmt.Sprintf("%.6f", float32(positive)/length),
		fmt.Sprintf("%.6f", float32(negative)/length),
		fmt.Sprintf("%.6f", float32(zeros)/length),
	}
}
