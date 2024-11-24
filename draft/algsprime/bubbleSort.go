package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 9, 7, 3, 2, 4}

	fmt.Println(len(s1))
	res := bubbleSearch(s1)

	fmt.Println(res)
}

func bubbleSearch(arr []int) []int {
	for i := range arr {
		for j := range len(arr) - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
