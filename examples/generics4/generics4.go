package main

import "fmt"

type Person struct {
	name string
	age  int
}

type CustomString string

func main() {
	Prt(CustomString("abc"))
	Prt(55)
}

type Stringer interface {
	~string | int
}

func Prt[T Stringer](str T) {
	fmt.Println(str)
}
