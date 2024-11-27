package main

import "fmt"

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
	queue := make([]*Node, 0)
	queue = append(queue, tree.root)

	fmt.Println("root:", tree.root.value)

	for len(queue) != 0 {
		node := queue[0]

		queue = append(queue[:0], queue[1:]...)
		if node.left != nil {
			queue = append(queue, node.left)
			fmt.Printf("parent: %v, left: %v ", node.value, node.left.value)
		}
		if node.right != nil {
			queue = append(queue, node.right)
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
	bst.insert(2)
	bst.insert(55)
	bst.insert(4)

	fmt.Printf("%+v\n", bst.root)

	bst.print()
}
