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

// Cell is a struct to store coordinates of specific grid.
type Cell struct {
	x         int64
	y         int64
	direction string
}

// nextCells finds the next open cell that is not visited.
func nextCells(grid [][]string, current Cell, visited map[Cell]bool, size int64) []Cell {
	next := make([]Cell, 0)
	next = searchX(grid, current, visited, size, next)
	next = searchY(grid, current, visited, size, next)

	return next
}

// searchX make a search over X grid cells.
func searchX(grid [][]string, current Cell, visited map[Cell]bool, size int64, next []Cell) []Cell {
	if current.direction == "" || current.direction == "Y" {
		x := current.x
		for i := x - 1; i >= 0; i-- {
			if grid[i][current.y] == "X" {
				break
			}

			cell := Cell{i, current.y, "X"}
			if _, ok := visited[cell]; !ok {
				next = append(next, cell)
			}
		}

		for i := x + 1; i < size; i++ {
			if grid[i][current.y] == "X" {
				break
			}

			cell := Cell{i, current.y, "X"}
			if _, ok := visited[cell]; !ok {
				next = append(next, cell)
			}
		}
	}

	return next
}

// searchY make a search over Y grid cells.
func searchY(grid [][]string, current Cell, visited map[Cell]bool, size int64, next []Cell) []Cell {
	if current.direction == "" || current.direction == "X" {
		y := current.y
		for i := y - 1; i >= 0; i-- {
			if grid[current.x][i] == "X" {
				break
			}

			cell := Cell{current.x, i, "Y"}
			if _, ok := visited[cell]; !ok {
				next = append(next, cell)
			}
		}

		for i := y + 1; i < size; i++ {
			if grid[current.x][i] == "X" {
				break
			}

			cell := Cell{current.x, i, "Y"}
			if _, ok := visited[cell]; !ok {
				next = append(next, cell)
			}
		}
	}

	return next
}

// minimumMoves find the minimum number of moves to reach the given goal from the departure point.
func minimumMoves(grid [][]string, startX int64, startY int64, goalX int64, goalY int64, visited map[Cell]bool, size int64) int64 {
	current := Cell{startX, startY, ""}
	target := Cell{goalX, goalY, ""}
	visited[current] = true

	queue := make([]Cell, 0)
	queue = append(queue, current)

	moves := make([]int64, 0)
	moves = append(moves, 0)

	for len(queue) > 0 {
		nextCell := queue[0]
		queue = queue[1:]

		move := moves[0]
		moves = moves[1:]

		if nextCell.x == target.x && nextCell.y == target.y {
			return move
		}

		next := nextCells(grid, nextCell, visited, size)
		for i := range next {
			if _, ok := visited[next[i]]; !ok {
				visited[next[i]] = true
			}
		}

		queue = append(queue, next...)

		for i := 0; i < len(next); i++ {
			moves = append(moves, move+1)
		}
	}

	return -1
}
