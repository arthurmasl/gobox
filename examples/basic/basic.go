package main

import (
	"fmt"
	"strings"
	"time"
)

type VectorI interface {
	Sum()
}

type Vector struct {
	x, y int
}

func (v *Vector) Sum() int {
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
		sum += vector.Sum()
	}

	fmt.Println(sum)
	fmt.Println(time.Since(t1))

	fmt.Println(items)

	elem, ok := items["three"]
	fmt.Println(elem, ok)

	fmt.Println(WordCount("hello world"))
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	for _, word := range strings.Fields(s) {
		res[string(word)] = 1
	}

	return res
}
