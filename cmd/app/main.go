package main

import "fmt"

type Set struct {
	data map[string]bool
}

func NewSet() *Set {
	return &Set{data: make(map[string]bool)}
}

func (set *Set) add(target string) {
	set.data[target] = true
}

func (set *Set) delete(target string) {
	delete(set.data, target)
}

func main() {
	fruits := NewSet()

	fruits.add("apple")
	fruits.add("apple")
	fruits.add("grape")
	fruits.delete("grape22")

	fmt.Println(len(fruits.data))

	for fruit := range fruits.data {
		fmt.Println(fruit)
	}
}
