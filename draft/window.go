package main

import (
	"fmt"
	"iter"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a2 := []string{"a", "b", "c", "d", "e"}

	for chunk := range Window(a1, 2) {
		fmt.Println(chunk)
	}

	for chunk := range Window(a2, 3) {
		fmt.Println(chunk)
	}
}

func Window[Slice ~[]E, E any](slice Slice, size int) iter.Seq[Slice] {
	return func(yield func(Slice) bool) {
		for i := range slice[:len(slice)-size+1] {
			if !yield(slice[i : i+size]) {
				return
			}
		}
	}
}
