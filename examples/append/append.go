package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	s1 := a1[:]

	fmt.Println(s1)

	fmt.Println(s1)

	r1 := append(append([]int{}, s1[:0]...), s1[1:]...)
	fmt.Println(r1)

	r2 := append(append([]int{}, s1[:1]...), s1[2:]...)
	fmt.Println(r2)
}
