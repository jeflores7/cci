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
	updateMap(node, depth, lLists)
	getLists(node.left, depth+1, lLists)
	getLists(node.right, depth+1, lLists)
	return
}

func updateMap(node *node, depth int, lLists map[int]*llNode) {
	llNode := &llNode{data: node}
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

//          10
//    1            5
// 0     2      6
//  -2    12     15

// lLists := {
// 0: 10,
// 1: 1 -> 5,
// 2: 0 -> 2 -> 6,
// 3: -2 -> 12 -> 15
// }

// depth := 3
// depthNumNodes := 8
// q := [nil nil nil nil nil nil nil nil nil nil nil nil nil nil]
// i := 6
// curNode := 15

// My implementation with weird tracking of depth and extra nils that are unnecessary
func getDepthListsBFS(root *node) map[int]*llNode {
	lLists := make(map[int]*llNode)
	if root == nil {
		return lLists
	}
	q := make([]*node, 1)
	q[0] = root
	depth := 0
	i := 0
	emptyDepth := false
	depthNumNodes := 1 // 2 ^ 0
	for len(q) > 0 {
		curNode := q[0]
		q = q[1:]
		i++
		addToMap(curNode, depth, lLists)
		if curNode == nil {
			q = append(q, nil, nil)
		} else {
			q = append(q, curNode.left, curNode.right)
		}
		if i == depthNumNodes {
			i = 0
			depthNumNodes = depthNumNodes * 2
			depth++
		}
		emptyDepth = true
		for _, v := range q {
			if v != nil {
				emptyDepth = false
				break
			}
		}
		if emptyDepth {
			break
		}
	}
	return lLists
}

func addToMap(node *node, depth int, lLists map[int]*llNode) {
	if node == nil {
		return
	}
	llNode := &llNode{data: node}
	if lLists[depth] == nil {
		lLists[depth] = llNode
		return
	}
	tail := lLists[depth]
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = llNode
}

func getDepthListsBFSBetter(root *node) map[int]*llNode {
	lLists := make(map[int]*llNode)
	if root == nil {
		return lLists
	}
	q := make([]*node, 1)
	q[0] = root
	depth := 0
	for len(q) > 0 {
		nextLevel := make([]*node, 0)
		for _, node := range q {
			updateMap(node, depth, lLists)
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		q = nextLevel
		depth++
	}
	return lLists
}
