package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{4, 2, 3, 4, 5}

	fmt.Println(len(arr), cap(arr), arr)
	fmt.Println(slices.Contains(arr, 6))

	arr = append(arr, 3)
	arr = append(arr, 1)
	slices.Sort(arr)
	fmt.Println(arr)
}
