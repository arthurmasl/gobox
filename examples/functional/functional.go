package main

import (
	"fmt"
	"strings"
)

type mapFunc[E any] func(E) E

func Map[S ~[]E, E any](s S, f mapFunc[E]) S {
	result := make(S, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

type keepFunc[E any] func(E) bool

func Filter[S ~[]E, E any](s S, f keepFunc[E]) S {
	result := S{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

type reduceFunc[E any] func(cur, next E) E

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []int{1, 2, 3}

	fmt.Println(s1)
	fmt.Println(Map(s1, strings.ToUpper))

	fmt.Println(Filter(s2, func(v int) bool {
		return v > 1
	}))

	sum := Reduce(s2, 0, func(curr, next int) int {
		return curr + next
	})

	fmt.Println(sum)
}
