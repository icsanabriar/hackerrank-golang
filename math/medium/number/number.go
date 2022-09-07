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

import "math"

// closer find the close number from the target given a and b.
func closer(target int32, a int32, b int32) int32 {
	d1 := math.Abs(float64(a - target))
	d2 := math.Abs(float64(b - target))

	if d1 <= d2 {
		return a
	}

	return b
}

// closestNumber find the multiple of x closest to a pow b.
func closestNumber(a int32, b int32, x int32) int32 {

	if a == 1 || b == 0 {
		return closer(1, 0, x)
	}

	if b < 0 {
		return 0
	}

	target := int32(1)

	for i := int32(0); i < b; i++ {
		target = target * a
	}

	return closer(target, target/x*x, (target/x+1)*x)
}
