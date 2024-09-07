package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	bits := 1_000_000
	bitArray := make([]byte, bits/8)

	var stringBuilder strings.Builder

	for i := 0; i < bits; i++ {
		bit := (bitArray[i/8] >> (i % 8)) & 1
		stringBuilder.WriteString(fmt.Sprintf("%d", bit))
	}

	bitString := stringBuilder.String()

	bitSize := unsafe.Sizeof(bitString) + uintptr(len(bitString))
	byteSize := bitSize / 8
	kbSize := byteSize / 1024
	mbSize := kbSize / 1024

	fmt.Printf("%d bits\n", bitSize)
	fmt.Printf("%d bytes\n", byteSize)
	fmt.Printf("%d kb\n", kbSize)
	fmt.Printf("%d mb\n", mbSize)
}
