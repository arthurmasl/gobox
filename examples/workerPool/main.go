package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := range numJobs {
		go worker(i, jobs, results)
	}

	for i := range numJobs {
		jobs <- i
	}
	close(jobs)

	total := 0
	for range numJobs {
		total += <-results
	}

	fmt.Println("result:", total)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for jobId := range jobs {
		fmt.Printf("worker %d started job %d\n", id, jobId)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, jobId)
		results <- 1
	}
}
