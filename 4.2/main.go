package main

// in := 0 1 2 3 4 5
// mid := 3
// node0.data := 3
// node0.left = node1
// node0.right = node2

// in1 := 0 1 2
// mid := 1
// node1.data := 1
// node1.left = node3
// node1.right = node4

// in3 := 0
// mid := 0
// node3.data = 0
// node3.left = node5
// node3.right = node6

// in5 := []
// node5 = nil

// in6 := []
// node6 = nil

// in4 := 2
// mid := 0
// node4.data = 2
// node4.left = node7  // nil
// node4.right = node8 // nil

// in2 := 4 5
// mid := 1
// node2.data := 5
// node2.left = node9
// node2.right = node10 // nil

// in9 := 4
// mid := 0
// node9.data = 4
// node9.left = node11 // nil
// node9.right = node12 // nil

type node struct {
	data  int
	left  *node
	right *node
}

func minHeightBST(in []int) *node {
	if len(in) == 0 {
		return nil
	}
	mid := len(in) / 2
	node := &node{
		data:  in[mid],
		left:  minHeightBST(in[:mid]),
		right: minHeightBST(in[mid+1:]),
	}
	return node
}
