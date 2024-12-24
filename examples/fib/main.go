package main

import (
	"fmt"
	"time"
)

const size = 85

func main() {
	measure(fibConstant, "constant")
	measure(fibMemo, "memo")
	// measure(fibRecursive, "recursive")
}

func measure(fn func(int) int, msg string) {
	t1 := time.Now()
	fmt.Println(msg)
	fmt.Println(fn(size))
	fmt.Println(time.Since(t1))
	fmt.Println()
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibConstant(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func fibMemo(n int) int {
	memo := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		if _, exists := memo[n]; !exists {
			memo[n] = fib(n-1) + fib(n-2)
		}

		return memo[n]
	}

	return fib(n)
}
