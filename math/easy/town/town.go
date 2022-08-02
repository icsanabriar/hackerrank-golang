// Copyright 2021 Iván Camilo Sanabria
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

// connectingTowns calculates the total number of routes to connect the first to last town.
func connectingTowns(n int32, routes []int32) int32 {

	modulo := int32(1234567)
	total := int32(1)

	for i := int32(0); i < (n - 1); i++ {
		total = (total * routes[i]) % modulo
	}

	return total
}
