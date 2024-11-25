package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			ch <- "A"
		}
	}()

	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}
