package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Sscanf("a:15 b:22", "a:%s b:%s", &a, &b)
	fmt.Println(a, b)
}
