package main

import (
	"fmt"
	"strings"
)

func main() {
	size := 15
	for i := range size {
		fmt.Println(strings.Repeat(" ", size-i), strings.Repeat("#", i*2))
	}
	fmt.Println(strings.Repeat(" ", size-(size/3)+1), strings.Repeat("#", size/3))
}
