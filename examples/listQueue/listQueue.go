package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)

	first := queue.Front()
	queue.Remove(first)

	first2 := queue.Front()
	queue.Remove(first2)

	first3 := queue.Front()
	fmt.Println(first3)
	queue.Remove(first3)

	fmt.Println(queue.Front(), queue.Back())
}
