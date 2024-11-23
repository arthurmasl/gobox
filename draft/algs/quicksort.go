package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 1, 2, 55, 11, 4}

	fmt.Println(arr)
	fmt.Println(quicksort(arr))
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, right []int

	for _, item := range arr {
		if item < pivot {
			left = append(left, item)
		} else if item > pivot {
			right = append(right, item)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}
