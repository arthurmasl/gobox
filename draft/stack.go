package main

import "fmt"

type Stack struct {
	elements []string
}

type Stacker interface {
	Push()
	Pop()
	Peek()
	IsEmpty()
	Size()
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return value, true
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	value := s.elements[len(s.elements)-1]

	return value, true
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	cars := &Stack{}

	cars.Push("bnw")
	cars.Push("audi")
	cars.Pop()

	fmt.Println(cars.Size())
	fmt.Println(cars.Peek())

	fmt.Println(cars.elements)
}
