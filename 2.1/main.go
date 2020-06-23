package main

// Questions:
// What kind of data is stored in the list?
// What kind of list (single/double)?
// What constraints do the values of the list follow?
// Do we want a return value?

// Assumptions:
// Integers
// -1000000 <= values <= 1000000
// Single
// No return

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Insert(v int) *Node {
	next := n
	for next.Next != nil {
		next = next.Next
	}
	next.Next = &Node{
		Value: v,
		Next:  nil,
	}
	return next.Next
}

// Bug in original implementation.
// Missed the need for creating copies to go into the Hash Table
func removeDupes(head *Node) {
	var (
		last      *Node
		current   *Node
		hashTable []*Node
	)
	hashTable = make([]*Node, 1000000)

	if head == nil {
		return
	}

	headCopy := &Node{
		Value: head.Value,
	}
	hashTable[hash(head.Value)%len(hashTable)] = headCopy
	last = head
	current = head.Next
	for current != nil {
		idx := hash(current.Value) % len(hashTable)
		if hashTable[idx] == nil {
			currentCopy := &Node{
				Value: current.Value,
			}
			hashTable[idx] = currentCopy
			last = current
		} else if hashTable[idx].Value != current.Value && hashTable[idx].Next == nil {
			// Hash collision with distinct values, insert node
			currentCopy := &Node{
				Value: current.Value,
			}
			hashTable[idx].Next = currentCopy
			last = current
		} else if hashTable[idx].Value == current.Value || hashTable[idx].Next.Value == current.Value {
			// Take advantage of the fact our hash table will only ever contain 2 unique nodes per index
			// due to the value constraints
			last.Next = current.Next
		}
		current = current.Next
	}
}

func hash(value int) int {
	if value < 0 {
		return -1 * value
	}
	return value
}

func removeDupesNoExtraDS(head *Node) {
	last := head
	outer := head
	for outer != nil {
		inner := outer.Next
		for inner != nil {
			if outer.Value == inner.Value {
				// Remove dupe
				last.Next = inner.Next
			} else {
				last = inner
			}
			inner = inner.Next
		}
		outer = outer.Next
		last = outer
	}
}
