package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func (c *Container) incrementBy(name string, n int) {
	fmt.Printf("increment %v by %d \n", name, n)
	time.Sleep(1000 * time.Millisecond)

	for i := 0; i < n; i++ {
		c.increment(name)
	}
	wg.Done()
}

func main() {
	c := Container{
		counters: map[string]int{"likes": 0, "dislikes": 0},
	}

	wg.Add(3)
	go c.incrementBy("likes", 1000)
	go c.incrementBy("likes", 500)
	go c.incrementBy("dislikes", 200)

	wg.Wait()
	fmt.Printf("%#v \n", c.counters)
}
