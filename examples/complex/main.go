package main

import (
	"fmt"
	"math"
)

func main() {
	grid := []string{
		"############",
		"#..........#",
		"#..........#",
		"#..........#",
		"#......E...#",
		"#..........#",
		"############",
	}

	a := complex(3, 1)
	b := complex(6, 2)

	fmt.Println(b - a)

	for y, row := range grid {
		newRow := []byte(row)

		for x := range row {
			if y == int(imag(a)) && x == int(real(a)) {
				newRow[x] = 'A'
			}
			if y == int(imag(b)) && x == int(real(b)) {
				newRow[x] = 'B'
			}
		}

		grid[y] = string(newRow)
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	point := complex(1, 0)
	rotated := rotate(point, 90)
	fmt.Println(point)
	fmt.Println(rotated)

	fmt.Printf("Original: %.2f + %.2fi, Rotated: %.2f + %.2fi\n",
		real(point), imag(point), real(rotated), imag(rotated))

	fmt.Println(math.Round(real(rotated)), math.Round(imag(rotated)))
}

func rotate(point complex128, angle float64) complex128 {
	theta := angle * math.Pi / 180
	rotation := complex(math.Cos(theta), math.Sin(theta))
	return point * rotation
}
