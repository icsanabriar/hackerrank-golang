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

// bfs estimates the shortest distance between the nodes given a starting one.
func bfs(n int, m int, edges [][]int64, s int64) []int64 {
	graph := make([][]int64, n+1)
	for i := 0; i < m; i++ {
		u := edges[i][0]
		v := edges[i][1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	distances := make([]int64, n+1)
	for i := range distances {
		distances[i] = -1
	}

	distances[s] = 0
	queue := []int64{s}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			if distances[v] == -1 {
				distances[v] = distances[u] + 6
				queue = append(queue, v)
			}
		}
	}

	distances = append(distances[:s], distances[s+1:]...)
	return distances[1:]
}
