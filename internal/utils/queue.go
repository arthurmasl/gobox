package utils

type queue[T any] struct {
	head, tail *node[T]
	Length     int
}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Insert(value T) {
	node := &node[T]{value: value}

	if q.tail == nil {
		q.head = node
		q.tail = q.head
		q.Length += 1
		return
	}

	q.tail.next = node
	q.tail = node
	q.Length += 1
}

func (q *queue[T]) Remove() T {
	var emptyValue T
	head := q.head

	if q.Length == 1 {
		q.head, q.tail = nil, nil
		q.Length = 0
		return head.value
	}

	if q.head == nil {
		return emptyValue
	}

	q.head = q.head.next
	q.Length -= 1

	return head.value
}
