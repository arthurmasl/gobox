package main

import (
	"fmt"
	"time"
)

func debounce(f func(), delay time.Duration) func() {
	var timer *time.Timer

	return func() {
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(delay, f)
	}
}

func main() {
	debouncedFunc := debounce(func() {
		fmt.Println("Debounced call")
	}, 1000*time.Millisecond)

	debouncedFunc()
	time.Sleep(time.Second * 2)
}
