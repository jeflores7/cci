package main

import "strings"

type Node struct {
	data rune
	next *Node
}

func (n *Node) Insert(r rune) *Node {
	for n.next != nil {
		n = n.next
	}
	n.next = &Node{
		data: r,
	}
	return n.next
}

// Questions:
// 1. Are spaces and punctuation allowed?
// 2. Are they counted or ignored if yes?
// 3. Is the palindrome testing case sensitivity?
// 4. Are we using ascii chars only?

// Assumptions:
// ASCII alphabet
// Spaces allowed, but ignored
// -- Removed tests with spaces due to recursive function request
// -- This implementation will handle spaces, but recursive won't
// Case insensitive check
func isPalindrome(head *Node) bool {
	if head == nil {
		return false
	}
	runner := head
	reverseList := &Node{data: head.data}
	numNodes := 1
	for runner.next != nil {
		runner = runner.next
		tmp := &Node{data: runner.data, next: reverseList}
		reverseList = tmp
		numNodes++
	}
	runner = head
	reverseRunner := reverseList
	for i := 0; i < numNodes/2; i++ {
		if strings.ToLower(string(runner.data)) != strings.ToLower(string(reverseRunner.data)) {
			return false
		}
		runner = runner.next
		for runner.data == ' ' {
			runner = runner.next
		}
		reverseRunner = reverseRunner.next
		for reverseRunner.data == ' ' {
			reverseRunner = reverseRunner.next
		}
	}
	return true
}

// Doesn't handle spaces in the list
func isPalindromeRecursive(curNode *Node, curPos int, length int) (*Node, bool) {
	if curNode == nil {
		return nil, false
	}
	// Int division will function like floor(length/2)
	if curPos == length/2 {
		if length%2 == 0 {
			// Even number of nodes, need to do a check against the next node
			return curNode.next.next, strings.ToLower(string(curNode.data)) == strings.ToLower(string(curNode.next.data))
		}
		// Odd number of nodes, the palindrome will share this character
		return curNode.next, true
	}
	next, ok := isPalindromeRecursive(curNode.next, curPos+1, length)
	if !ok {
		return nil, false
	}
	if next == nil {
		return nil, false
	}
	return next.next, strings.ToLower(string(curNode.data)) == strings.ToLower(string(next.data))
}
