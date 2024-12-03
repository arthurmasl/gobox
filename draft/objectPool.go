package main

import (
	"fmt"

	"gobox/internal/utils"
)

type ObjectPool struct {
	data  []int
	index int
}

func NewObjectPool(size int) *ObjectPool {
	return &ObjectPool{
		data: make([]int, size),
	}
}

func (pool *ObjectPool) Insert(value int) {
	pool.data[pool.index] = value
	pool.index = (pool.index + 1) % len(pool.data)
}

func main() {
	op := NewObjectPool(4)

	op.Insert(1)
	op.Insert(2)
	op.Insert(3)
	op.Insert(4)
	utils.Assert(op.data[0] == 1)
	op.Insert(5)
	utils.Assert(op.data[0] == 5)

	fmt.Println(op.data, op.index)
}
