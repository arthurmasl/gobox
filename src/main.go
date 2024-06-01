package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello22, world")

	nums := []int{1, 2, 3, 4, 5}

	for n, v := range nums {
		fmt.Println(n, v)
	}
}
