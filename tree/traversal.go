package main

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func (node *Node) TraversalFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversalFunc(f)
	f(node)
	node.Right.TraversalFunc(f)
}

func (node *Node) Traverse() {
	node.TraversalFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func main() {
	var root Node

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = &Node{2, nil, nil}

	root.Traverse()
	nodeCount := 0
	root.TraversalFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("tree node count is ", nodeCount)
}
