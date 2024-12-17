package main

import "fmt"

func main() {
	input := []int{5, 7, 1, 4, 9, 3, 15}
	var bitmap uint16

	for i := range 16 {
		bitmap |= (0 << i)
	}

	for _, v := range input {
		bitmap |= (1 << v)
	}

	result := make([]int, 16)
	resultIndex := 0
	for i := range 16 {
		if (bitmap & (1 << i)) != 0 {
			result[resultIndex] = i
			resultIndex++
		}
	}

	fmt.Println(input)
	fmt.Printf("%016b\n", bitmap)
	fmt.Println(result)
}
