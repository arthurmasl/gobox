package main

import (
	"fmt"
	"gobox/internal/utils"
	"strings"
)

type BinaryTree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (tree *BinaryTree) insert(value int) {
	node := &Node{value: value}
	current := tree.root

	for {
		if current == nil {
			current = node
			break
		}

		if current.value > node.value {
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

func (tree *BinaryTree) print() {
	queue := utils.NewQueue[*Node]()
	queue.Insert(tree.root)

	items := 0
	expected := 1
	spaces := 0

	for queue.Length != 0 {
		node := queue.Remove()
		if node == nil {
			continue
		}

		if expected == 1 {
			spaces = 16
		}
		fmt.Printf("%v%v", strings.Repeat(" ", spaces), node.value)
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

		if node.left != nil {
			queue.Insert(node.left)
		}

		if node.right != nil {
			queue.Insert(node.right)
		}

	}
}

func main() {
	bst := BinaryTree{}
	bst.root = &Node{value: 20}

	bst.insert(15)
	bst.insert(25)

	bst.insert(14)
	bst.insert(16)
	bst.insert(24)
	bst.insert(26)

	bst.insert(13)
	bst.insert(17)
	bst.insert(15)
	bst.insert(18)
	bst.insert(23)
	bst.insert(27)
	bst.insert(22)
	bst.insert(30)

	bst.print()
	fmt.Println()
}
