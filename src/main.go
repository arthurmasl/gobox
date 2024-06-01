package main

import (
	"fmt"
	"os"
	"slices"
)

type Vector struct {
	x, y, z int
}

func (v *Vector) sum() int {
	return v.x + v.y + v.z
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	v1 := Vector{10, 20, 1}
	v1.z = 55

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
