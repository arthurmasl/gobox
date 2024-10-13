package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{"bob", 22}
	p2 := Person{"bob", 22}

	s1 := []Person{p1}
	s2 := []Person{p2}

	fmt.Println(Insert(s1, s2))
}

func Insert[S ~[]E, E comparable](a, b S) S {
	return append(a, b...)
}
