package main

//          10
//    1            5
// 0     2      6
//  -2    12     15

// lLists := {
// 0: node10,
// 1: node1 -> node5,
// 2: node0 -> node2 -> node6,
// 3: node-2 -> node12 -> node15,
// }

// depth := 3
// node := 15

type node struct {
	data  int
	left  *node
	right *node
}

type llNode struct {
	data *node
	next *llNode
}

func getDepthLists(root *node) map[int]*llNode {
	lLists := make(map[int]*llNode)
	getLists(root, 0, lLists)
	return lLists
}

func getLists(node *node, depth int, lLists map[int]*llNode) {
	if node == nil {
		return
	}
	llNode := &llNode{data: node}
	updateMap(llNode, depth, lLists)
	getLists(node.left, depth+1, lLists)
	getLists(node.right, depth+1, lLists)
	return
}

func updateMap(llNode *llNode, depth int, lLists map[int]*llNode) {
	if lLists[depth] == nil {
		lLists[depth] = llNode
		return
	}
	lListNode := lLists[depth]
	for lListNode.next != nil {
		lListNode = lListNode.next
	}
	lListNode.next = llNode
	return
}
