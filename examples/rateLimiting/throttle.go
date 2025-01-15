package main

import (
	"fmt"
	"time"
)

func main() {
	const numRequests = 15

	requests := make(chan int, numRequests)
	for i := range numRequests {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1000 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
