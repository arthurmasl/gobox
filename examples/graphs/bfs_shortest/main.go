package main

import (
	"container/list"
	"fmt"

	"gobox/internal/utils"
)

type vec struct {
	x, y int
}

var directions = []vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

var grid = []string{
	"S..#...",
	"..#..#.",
	"....#..",
	"...#..#",
	"..#..#.",
	".#..#..",
	"#.#...E",
}

func main() {
	start, end := getInitialPositions()

	// store parents to backtrack path later
	// instead of this can be used distance
	// distance := make(map[Vector]int)
	parent := make(map[vec]vec)

	// store path to draw later
	path := make([]vec, 0)

	visited := make(map[vec]bool)
	visited[start] = true

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(vec)
		queue.Remove(node)

		if current == end {
			path = make([]vec, 0)
			for current != start {
				path = append(path, current)
				current = parent[current]
			}

			fmt.Println("found shortest path, distance:", len(path))
			break
		}

		for _, neighbor := range getNeighbors(grid, current) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	for _, pos := range path {
		if grid[pos.y][pos.x] != 'E' {
			newRow := []byte(grid[pos.y])
			newRow[pos.x] = 'o'
			grid[pos.y] = string(newRow)
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}

func getNeighbors(grid []string, pos vec) []vec {
	neighbors := make([]vec, 0)

	for _, dir := range directions {
		neighborPos := vec{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(grid, neighborPos.x, neighborPos.y)

		if ok && neighborId != '#' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}

func getInitialPositions() (vec, vec) {
	var start, end vec
	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				start = vec{x, y}
			}
			if col == 'E' {
				end = vec{x, y}
			}
		}
	}

	return start, end
}
