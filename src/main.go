package main

import (
	"fmt"
	"time"
)

type Vector struct {
	x, y int
}

func (v *Vector) sum() int {
	return v.x + v.y
}

var vectors [10000000]Vector

var items = map[string]Vector{
	"one": {10, 10},
	"two": {20, 20},
}

func main() {
	t1 := time.Now()
	sum := 0

	for i, vector := range vectors {
		vector.x = i
		sum += vector.sum()
	}

	fmt.Println(sum)
	fmt.Println(time.Since(t1))

	fmt.Println(items)
}
