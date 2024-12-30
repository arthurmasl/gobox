package main

import (
	"container/list"
	"fmt"

	"gobox/internal/utils"
)

var directions = []complex128{complex(0, -1), complex(1, 0), complex(0, 1), complex(-1, 0)}

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
	parent := make(map[complex128]complex128)

	// store path to draw later
	path := make([]complex128, 0)

	visited := make(map[complex128]bool)
	visited[start] = true

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(complex128)
		queue.Remove(node)

		if current == end {
			path = make([]complex128, 0)
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
		x, y := int(real(pos)), int(imag(pos))
		if grid[y][x] != 'E' {
			newRow := []byte(grid[y])
			newRow[x] = 'o'
			grid[y] = string(newRow)
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}

func getNeighbors(grid []string, pos complex128) []complex128 {
	neighbors := make([]complex128, 0)

	for _, dir := range directions {
		neighborPos := pos + dir
		x, y := getVec(neighborPos)
		neighborId, ok := utils.GetSafeValue(grid, x, y)

		if ok && neighborId != '#' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}

func getInitialPositions() (complex128, complex128) {
	var start, end complex128
	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				start = newComplex(x, y)
			}
			if col == 'E' {
				end = newComplex(x, y)
			}
		}
	}

	return start, end
}

func getVec(c complex128) (int, int) {
	return int(real(c)), int(imag(c))
}

func newComplex(x, y int) complex128 {
	return complex(float64(x), float64(y))
}
