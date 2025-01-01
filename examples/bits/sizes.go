package main

import (
	"fmt"
	"math"
)

func main() {
	const big = 1<<8 - 1
	fmt.Println(uint8(1<<8 - 1))
	fmt.Println(uint16(1<<16 - 1))
	fmt.Println(uint32(1<<32 - 1))
	fmt.Println(uint64(1<<64 - 1))

	fmt.Println(int8(1<<7 - 1))

	fmt.Println(^uint32(0))
	fmt.Println(math.MaxUint32)

	fmt.Println((1<<32 - 1) >> 8)

	for i := range 33 {
		fmt.Printf("%032b%032b\n", 1<<i-1, (1<<32-1)>>i)
	}
	for i := range 33 {
		fmt.Printf("%032b%032b\n", (1<<32-1)>>i, 1<<i-1)
	}

	fmt.Println()
	for i := range 9 {
		fmt.Printf("%08b\n", 1<<i-1)
	}
}
