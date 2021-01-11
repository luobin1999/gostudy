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

func (node *Node) TraverseWithChannel() chan *Node {
	channel := make(chan *Node)
	go func() {
		node.TraversalFunc(func(node *Node) {
			channel <- node
		})
		close(channel)
	}()
	return channel
}

func main() {
	var root Node

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{10, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = &Node{2, nil, nil}

	root.Traverse()
	nodeCount := 0
	root.TraversalFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("tree node count is ", nodeCount)

	c := root.TraverseWithChannel()
	maxValue := 0
	for n := range c {
		if n.Value > maxValue {
			maxValue = n.Value
		}
	}
	fmt.Println("max node value is ", maxValue)
}
