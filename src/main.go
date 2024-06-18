package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(rand.IntN(100))
	fmt.Println(rand.NewPCG(5, 10).Uint64())
}
