package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := make(chan time.Time, 3)

	for range 3 {
		limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			limiter <- t
		}
	}()

	requests := make(chan int, 15)
	for i := range 15 {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
