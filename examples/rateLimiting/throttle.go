package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := range 5 {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1000 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
