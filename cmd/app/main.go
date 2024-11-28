package main

import (
	"fmt"
	"strings"

	"gobox/internal/utils"
)

type BinaryTree struct {
	root  *Node
	queue *utils.Queue[*Node]
}

type Node struct {
	Value int
	left  *Node
	right *Node
}

func (tree *BinaryTree) Traverse() bool {
	if tree.queue == nil {
		tree.queue = utils.NewQueue[*Node]()
		tree.queue.Insert(tree.root)
	}

	return tree.queue.Length > 0
}

func (tree *BinaryTree) Next() *Node {
	queue := tree.queue
	node := queue.Remove()

	if node.left != nil {
		queue.Insert(node.left)
	}

	if node.right != nil {
		queue.Insert(node.right)
	}

	return node
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

		if node.left != nil {
			queue.Insert(node.left)
		}

		if node.right != nil {
			queue.Insert(node.right)
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

	for tree.Traverse() {
		node := tree.Next()
		fmt.Println(node.Value)
	}

	tree.Print()
}
