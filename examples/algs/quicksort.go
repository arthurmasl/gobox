package main

import (
	"math/rand/v2"

	"gobox/internal/utils"
)

func main() {
	arrSize := 10_000_000
	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = rand.IntN(arrSize)
	}

	defer utils.Perf()()
	// 550ms
	okSort(arr, 0, len(arr)-1)

	// 1.5s
	// noobSort(arr)

	// 600ms
	// slices.Sort(arr)
}

func noobSort(arr []int) []int {
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

	return append(append(noobSort(left), pivot), noobSort(right)...)
}

func okSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		okSort(arr, low, p-1)
		okSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}
