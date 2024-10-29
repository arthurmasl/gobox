package main

import (
	"fmt"
)

func main() {
	b1 := uint8(0b00000100)

	fmt.Printf("toggle %08b\n", b1^(1<<2))

	// set
	fmt.Printf("set %08b\n", b1|(1<<0))

	// clear
	fmt.Printf("clear %08b\n", b1&^(1<<2))

	// check
	fmt.Println("check", (b1&(1<<2)) != 0)
}
