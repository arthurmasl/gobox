package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	total := 0

	for range 3 * 1000000000 {
		total += 1
	}

	fmt.Println(total)
	fmt.Println(time.Since(t1))
}
