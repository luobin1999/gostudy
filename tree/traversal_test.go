package main

func ExampleNode_Print() {
	node1 := Node{1, nil, nil}
	node2 := Node{5, nil, nil}
	node1.Print()
	node2.Print()
	// Output
	// 1
	// 5
}
