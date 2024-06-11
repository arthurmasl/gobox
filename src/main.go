package main

import (
	"fmt"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

func main() {
	m := map[int]string{1: "1", 2: "2"}

	fmt.Printf("%#v", MapKeys(m))
}
