package main

import (
	"fmt"
	"os"
	"slices"
)

type Vector struct {
	x, y int
}

func (v *Vector) sum() int {
	return v.x + v.y
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	v1 := Vector{10, 20}
	v1.x = 55

	fmt.Println(v1)
	fmt.Println(v1.sum())

	text := []byte("hello, world")
	e := os.WriteFile("test.txt", text, 0)
	check(e)

	strings := []string{"c", "a", "b"}
	nums := []int{5, 54, 1, 13, 65}

	slices.Sort(strings)
	slices.Sort(nums)

	fmt.Println(strings)
	fmt.Println(nums)
}
