package main

import (
	"fmt"
	"iter"
)

func Traverse() iter.Seq[int] {
	return func(yield func(int) bool) {
		items := []int{1, 2, 3, 4}

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

func main() {
	for item := range Traverse() {
		fmt.Println(item)
	}
}
