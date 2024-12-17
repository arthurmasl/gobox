package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	c := make(chan int)

	for i := range 3 {
		fmt.Println("start channel", i)
		go calc(1000000000, c)
	}

	c1, c2, c3 := <-c, <-c, <-c

	fmt.Println(c1 + c2 + c3)
	fmt.Println(time.Since(t1))
}

func calc(n int, c chan int) {
	sum := 0

	for range n {
		sum += 1
	}

	c <- sum
}
