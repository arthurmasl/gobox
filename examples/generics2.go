package main

import "fmt"

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T

	for _, v := range slice {
		sum += v
	}

	return sum
}

func main() {
	intSlice := []int{1, 2, 3}
	float32Slice := []float32{1, 2, 3.5}

	fmt.Println(sumSlice(intSlice))
	fmt.Println(sumSlice(float32Slice))
}
