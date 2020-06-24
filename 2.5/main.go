package main

type Node struct {
	data int64
	next *Node
}

// Questions:
// 1. How big can our linked list be? (This would tell us if we're dealing with a specific type of int (32, 64) or if we need to use math package
// Assume int64 is large enough
// 2. Should we default to returning a node with data == 0?
// Sure
// 3. Are the lists guaranteed to both be equal length?
// Assume no
func sumReverseLinkedLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil || list2 == nil {
		return &Node{data: 0}
	}
	v1 := list1
	v1Length := 0
	for v1 != nil {
		v1Length++
		v1 = v1.next
	}
	v2 := list2
	v2Length := 0
	for v2 != nil {
		v2Length++
		v2 = v2.next
	}
	big := list1
	small := list2
	if v2Length > v1Length {
		big = list2
		small = list1
	}
	// When writing in Docs forgot to cast to int64 for compatible operations with Node.data
	multiplier := int64(1)
	sumData := int64(0)
	var sum, sumHead *Node
	for big != nil {
		currentBig := big.data * multiplier
		if small == nil {
			sumData += currentBig
		} else {
			sumData += currentBig + (small.data * multiplier)
			small = small.next
		}
		big = big.next

		nodeData := (sumData / multiplier) % 10
		if sum == nil {
			sum = &Node{data: nodeData}
			// After updating, forgot to keep a pointer to sum head for return
			sumHead = sum
		} else {
			sum.next = &Node{data: nodeData}
			// Missed moving sum over as we add nodes
			sum = sum.next
		}
		multiplier = multiplier * 10
	}
	// When writing in docs had comparison the wrong direction
	for sumData > multiplier {
		nodeData := (sumData / multiplier) % 10
		sum.next = &Node{data: nodeData}
		multiplier = multiplier * 10
		sum = sum.next
	}
	return sumHead
}

// Goal: 21
// 9  1
//         |
// 2
//     |
// sumData: 21
// multiplier: 100

// Sum list
// 1  2