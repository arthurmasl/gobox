package main

import (
	"fmt"
	"time"
)

func main() {
	b1 := uint16(1)
	i1 := uint16(0)

	for {
		time.Sleep(time.Millisecond * 100)

		b1 = b1 << 1
		i1++

		if b1 == 0 {
			b1 = ^b1 >> 15
		}

		fmt.Print("\033[H\033[2J")

		fmt.Printf("%016b\n", b1)
		fmt.Println(b1)

		fmt.Printf("%016b\n", i1)
		fmt.Println(i1)
	}
}
