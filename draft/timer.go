package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Microsecond * 200)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	<-done
	fmt.Println("stop")
}
