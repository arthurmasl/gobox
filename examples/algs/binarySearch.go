package main

func main() {
	nums := []int{1, 2, 3, 5, 6, 7, 9, 10, 20, 25, 30, 33, 34}
	index := binarySearch(nums, 3)

	println(index)
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
