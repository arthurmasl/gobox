package main

import (
	"fmt"

	"gobox/internal/utils"
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
			fmt.Println("isert")
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

	fmt.Println("root:", tree.root.value)

	for queue.Length != 0 {
		node := queue.Remove()

		if node == nil {
			continue
		}

		if node.left != nil {
			queue.Insert(node.left)
			fmt.Printf("parent: %v, left: %v ", node.value, node.left.value)
		}
		if node.right != nil {
			queue.Insert(node.right)
			fmt.Printf("parent: %v, right: %v\n", node.value, node.right.value)
		}
	}
}

func main() {
	bst := BinaryTree{}
	bst.root = &Node{value: 10}
	bst.insert(15)
	bst.insert(5)
	bst.insert(25)
	bst.insert(7)

	// fmt.Printf("%+v\n", bst.root)
	bst.print()
}
