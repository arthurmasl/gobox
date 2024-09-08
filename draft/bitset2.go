package main

import (
	"fmt"
)

type BitSet uint8

type Bitseter interface {
	Set()
	Clear()
	Toggle()
}

type Bitbuger interface {
	Check()
	Print()
}

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

func debug(b Bitbuger) {
	b.Print()
	b.Check()
}

func main() {
	b := BitSet(0)

	b.Set(0)
	b.Set(1)
	b.Clear(1)
	b.Toggle(0)
	b.Toggle(0)
	fmt.Println(b.Check(0))

	b.Print()
}
