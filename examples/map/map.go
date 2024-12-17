package main

import (
	"fmt"
)

func mapSlice(slice []int, callback func(int) int) []int {
	newSlice := make([]int, len(slice))

	for k, v := range slice {
		newSlice[k] = callback(v)
	}
	return newSlice
}

func pow(n int) int {
	return n * 2
}

func main() {
	slice1 := []int{2, 3, 4}
	mappedSlice := mapSlice(slice1, pow)

	fmt.Println(slice1)
	fmt.Println(mappedSlice)
}
