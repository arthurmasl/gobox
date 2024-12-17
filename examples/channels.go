package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	CHANNELS    = 10
	NUMS_TO_SUM = 1000000000
)

func main() {
	t1 := time.Now()
	wg := sync.WaitGroup{}
	ch := make(chan int, CHANNELS)

	fmt.Printf("Start processing %v numbers on %v threads\n", NUMS_TO_SUM, CHANNELS)
	for range CHANNELS {
		wg.Add(1)
		workerNums := NUMS_TO_SUM / CHANNELS

		go worker(&wg, ch, workerNums)
	}

	wg.Wait()
	close(ch)

	total := 0
	for res := range ch {
		total += res
	}

	fmt.Printf("Done in %v, result is %v\n", time.Since(t1), total == NUMS_TO_SUM)
}

func worker(wg *sync.WaitGroup, c chan int, numsToWork int) {
	defer wg.Done()

	sum := 0
	for range numsToWork {
		sum += 1
	}

	c <- sum
}
