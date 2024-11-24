package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := make([]bool, 10)
	s1[7] = true
	s1[8] = true
	s1[9] = true

	fmt.Println(s1)

	res := twoCrystalBalls(s1)
	fmt.Println(res)
}

func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))

	i := jumpAmount
	j := 0

	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jumpAmount
	}

	i -= jumpAmount

	for j < int(jumpAmount) && i < len(breaks) {
		if breaks[i] {
			return i
		}

		j++
		i++
	}

	return -1
}
