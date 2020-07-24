package main

type node struct {
	data  int
	left  *node
	right *node
}

func isBST(root *node) bool {
	if root == nil {
		return true
	}
	return checkBST(root.left, nil, root) && checkBST(root.right, root, nil)
}

func checkBST(n *node, min *node, max *node) bool {
	if n == nil {
		return true
	}

	if min != nil && n.data <= min.data {
		return false
	}
	if max != nil && n.data > max.data {
		return false
	}

	return checkBST(n.left, min, n) && checkBST(n.right, n, max)
}
