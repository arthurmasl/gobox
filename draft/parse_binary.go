package main

import (
	"fmt"
	"strconv"
)

func main() {
	binary := "11111000"
	n, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("aa")
	}

	fmt.Println(n)
	// fmt.Println((n >> 6) & 1)
	// fmt.Println((n >> 6))
	fmt.Println(strconv.FormatInt(int64(n>>2), 2))
}
