package main

type node struct {
	data   int
	parent *node
	left   *node
	right  *node
}

func findNext(n *node) *node {
	if n == nil {
		return nil
	}
	if n.right != nil {
		return findMin(n.right)
	}
	return findNextAncestor(n.parent, n)
}

func findMin(n *node) *node {
	if n.left == nil {
		return n
	}
	return findMin(n.left)
}

func findNextAncestor(parent *node, n *node) *node {
	if parent == nil || parent.data >= n.data {
		return parent
	}
	return findNextAncestor(parent.parent, n)
}
