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

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func reverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))

	for i, v := range slice {
		reversed[len(slice)-1-i] = v
	}

	return reversed
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	mappedSlice := mapSlice(slice1, pow)
	reverseSlice := reverseSlice(slice1)

	fmt.Println(slice1)
	fmt.Println(mappedSlice)
	fmt.Println(reverseSlice)
}
