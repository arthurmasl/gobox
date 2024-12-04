package utils

import (
	"container/list"
	"fmt"
)

type Stack[T any] struct {
	data *list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: list.New()}
}

func (stack *Stack[T]) Push(value T) {
	stack.data.PushBack(value)
}

func (stack *Stack[T]) Pop() *list.Element {
	if stack.data.Len() == 0 {
		return nil
	}

	tail := stack.data.Back()
	stack.data.Remove(tail)
	return tail
}

func main() {
	stack := NewStack[int]()

	stack.Push(1)
	fmt.Println(stack.Pop().Value)
	fmt.Println(stack.Pop())
}
