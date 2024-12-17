package main

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	nums := make([]int, 30)
	for i := range nums {
		nums[i] = i
	}

	target := 3
	result := binarySearch(nums, target)

	fmt.Println(result)

	if result != target {
		t.Errorf("test failed")
	}
}
