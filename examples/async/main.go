package main

import (
	"fmt"
	"time"
)

func main() {
	task1 := asyncTask(5)
	task2 := asyncTask(10)

	fmt.Println("other")

	r1 := <-task1
	fmt.Println(r1)

	r2 := <-task2
	fmt.Println(r2)
}

func asyncTask(input int) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		time.Sleep(time.Millisecond * time.Duration(input) * 100)
		resultChan <- input + 222
	}()

	return resultChan
}
