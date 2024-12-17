package main

import "fmt"

type Set struct {
	elements map[string]bool
}

func NewSet() *Set {
	return &Set{elements: make(map[string]bool)}
}

func (set *Set) add(target string) {
	set.elements[target] = true
}

func (set *Set) remove(target string) {
	delete(set.elements, target)
}

func main() {
	fruits := NewSet()

	fruits.add("apple")
	fruits.add("apple")
	fruits.add("grape")
	fruits.remove("grape22")

	fmt.Println(len(fruits.elements))

	for fruit := range fruits.elements {
		fmt.Println(fruit)
	}
}
