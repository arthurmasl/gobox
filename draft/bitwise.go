package main

import (
	"fmt"
	"unsafe"
)

var (
	READONLY = 0x1
	HIDDEN   = 0x2
	ARCHIVE  = 0x10
)

func getBitValue(num, n int) int {
	return (num >> n) & 1
}

func main() {
	n := uint8(42) // 101010
	numbits := int(unsafe.Sizeof(n) * 8)

	fmt.Println(numbits)
	fmt.Println("-----")

	for i := 0; i < numbits; i++ {
		fmt.Print(getBitValue(int(n), numbits-i-1))
	}
	fmt.Println()

	code := 0x1

	if (code & READONLY) != 0 {
		fmt.Println("readonly")
	}
	if (code & HIDDEN) != 0 {
		fmt.Println("hidden")
	}
	if (code & ARCHIVE) != 0 {
		fmt.Println("archive")
	}
}
