package main

import (
	"fmt"
	"iter"
	"strings"

	"gobox/internal/utils"
)

type BinaryTree struct {
	root *Node
}

type Node struct {
	Value int
	left  *Node
	right *Node
}

func (tree *BinaryTree) Traverse() iter.Seq[*Node] {
	queue := utils.NewQueue[*Node]()
	queue.Insert(tree.root)

	return func(yield func(*Node) bool) {
		for queue.Length != 0 {
			node := queue.Remove()
			if !yield(node) {
				return
			}

			if node.left != nil {
				queue.Insert(node.left)
			}
			if node.right != nil {
				queue.Insert(node.right)
			}
		}
	}
}

func (tree *BinaryTree) Insert(value int) {
	node := &Node{Value: value}
	current := tree.root

	for current != nil {
		if current.Value > node.Value {
			if current.left == nil {
				current.left = node
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = node
				break
			}
			current = current.right
		}
	}
}

func (tree *BinaryTree) Print() {
	items := 0
	expected := 1
	spaces := 0

	for node := range tree.Traverse() {
		if node == nil {
			continue
		}

		if expected == 1 {
			spaces = 16
		}
		fmt.Printf("%v%v", strings.Repeat(" ", spaces), node.Value)
		if expected == 1 {
			spaces = 20
		}
		items++

		if items == expected {
			fmt.Println()
			items = 0
			spaces /= 2

			if expected == 1 {
				expected = 2
			} else {
				expected *= 2
			}
		}

	}
}

func main() {
	tree := BinaryTree{root: &Node{Value: 20}}

	tree.Insert(15)
	tree.Insert(25)

	tree.Insert(14)
	tree.Insert(16)
	tree.Insert(24)
	tree.Insert(26)

	tree.Insert(13)
	tree.Insert(17)
	tree.Insert(15)
	tree.Insert(18)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(22)
	tree.Insert(30)

	tree.Print()

	for node := range tree.Traverse() {
		fmt.Println(node.Value)
	}
}
