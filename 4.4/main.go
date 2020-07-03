package main

//                                 n
//                    a                          b
//              c                        d               e
//        f                          g                          h
//                                       i

type node struct {
	left  *node
	right *node
}

func isBalanced(root *node) bool {
	if root == nil {
		return false
	}
	q := make([]*node, 1)
	q[0] = root
	currentDepth := 0
	nilChildNodeDepth := -1
	for len(q) > 0 {
		nextDepth := make([]*node, 0)
		for i := 0; i < len(q); i++ {
			if q[i].left != nil {
				nextDepth = append(nextDepth, q[i].left)
			} else if nilChildNodeDepth == -1 {
				nilChildNodeDepth = currentDepth
			}
			if q[i].right != nil {
				nextDepth = append(nextDepth, q[i].right)
			} else if nilChildNodeDepth == -1 {
				nilChildNodeDepth = currentDepth
			}
		}
		currentDepth++
		q = nextDepth
		if nilChildNodeDepth != -1 &&
			len(q) > 0 &&
			currentDepth-nilChildNodeDepth > 1 {
			return false
		}
	}
	return true
}

// currentDepth := 3, nilChildNodeDepth := 1
// q[i]      :=
// q         := [f g h]
// nextDepth := [f g h]
