package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for id := range 4 {
		go reader(id, ch, &wg)
	}

	for i := range 12 {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func reader(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	for {
		val, ok := <-ch
		if !ok {
			fmt.Printf("channel %v is closed\n", id)
			return
		}

		fmt.Printf("Reader %v processign channel %v\n", id, val)
	}
}
