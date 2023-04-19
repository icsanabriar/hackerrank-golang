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

// module is the value to simplify the result.
const module = 1000000007

// summingSeries calculates the sum of specific series when the nth term is n2 - (n-1)2.
// Simplifying the expression it is the sum of nth odd numbers which is equal a n2.
func summingSeries(n int64) int64 {

	n %= module

	return n * n % module
}
