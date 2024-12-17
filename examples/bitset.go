package main

import (
	"fmt"
	"strconv"
)

type BitSet uint8

func (b *BitSet) Set(pos uint) {
	*b |= (1 << pos)
}

func (b *BitSet) Clear(pos uint) {
	*b &= ^(1 << pos)
}

func (b *BitSet) Toggle(pos uint) {
	*b ^= (1 << pos)
}

func (b BitSet) Check(pos uint) bool {
	return (b & (1 << pos)) != 0
}

func (b BitSet) Print() {
	fmt.Printf("%08b\n", b)
}

func main() {
	// b := BitSet(0)
	//
	// b.Set(0)
	// b.Set(1)
	// b.Clear(1)
	// b.Toggle(0)
	// b.Toggle(0)
	// fmt.Println(b.Check(0))
	//
	// b.Print()

	binary := "00000100"
	n, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("aa")
	}

	fmt.Printf("%08b - initial\n", n)
	// fmt.Printf("%08b - mask\n", (1 << 2))
	fmt.Printf("%08b - mask xor\n", ^(1 << 2))
	n &= ^(1 << 2)
	fmt.Printf("%08b - and\n", n)
}
