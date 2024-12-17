package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("done all")
}

func worker(n int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d started\n", n)
	defer wg.Done()

	time.Sleep(time.Millisecond * 500)
	fmt.Printf("worker %d done\n", n)
}
