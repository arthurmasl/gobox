package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5
	results := make(chan int, numJobs)

	for i := range numJobs {
		go worker(i, results)
	}

	total := 0
	for range numJobs {
		total += <-results
	}

	fmt.Println("result:", total)
}

func worker(id int, results chan<- int) {
	fmt.Printf("worker %d started job\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished job\n", id)
	results <- 1
}
