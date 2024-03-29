// Copyright 2021 Iván Camilo Sanabria
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

// cutPaper calculates how many cuts are required taking into account the size of the paper on the parameters n x m.
func cutPaper(n int64, m int64) int64 {
	return (n-1)*(m-1) + (n - 1) + (m - 1)
}
