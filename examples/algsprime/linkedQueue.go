package main

import "fmt"

type Queue struct {
	head, tail *Node
	length     int
}

type Node struct {
	value int
	next  *Node
}

func (q *Queue) insert(value int) {
	node := &Node{value: value}

	if q.tail == nil {
		q.head = node
		q.tail = q.head
		q.length += 1
		return
	}

	q.tail.next = node
	q.tail = node
	q.length += 1
}

func (q *Queue) remove() int {
	if q.length == 1 {
		q.head, q.tail = nil, nil
		q.length = 0
		return -1
	}

	if q.head == nil {
		return -1
	}

	head := q.head
	q.head = q.head.next
	q.length -= 1

	return head.value
}

func main() {
	q := Queue{}

	q.insert(1)
	q.insert(2)
	q.insert(3)

	q.remove()
	q.remove()
	q.remove()

	fmt.Println(q.head, q.tail, q.length)
}
