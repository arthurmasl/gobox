package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 30)

	for i := range nums {
		nums[i] = i
	}

	fmt.Println(binarySearch(nums, 3))
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return arr[mid]
		}

		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}
