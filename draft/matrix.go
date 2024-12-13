package main

import (
	"fmt"
	"slices"
)

func main() {
	matrix := createMatrix(3)

	printMatrix(matrix)
	rotate(matrix)
	printMatrix(matrix)
}

func createMatrix(size int) [][]int {
	i := 1
	matrix := make([][]int, size)

	for y := range matrix {
		matrix[y] = make([]int, size)

		for x := range matrix[y] {
			matrix[y][x] = i
			i++
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println()
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := range n {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		slices.Reverse(row)
	}
}
