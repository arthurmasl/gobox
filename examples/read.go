package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}
