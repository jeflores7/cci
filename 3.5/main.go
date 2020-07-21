package main

import "errors"

type stack struct {
	elements []int
}

func NewStack() stack {
	return stack{
		elements: make([]int, 0),
	}
}

func (s *stack) push(el int) {
	s.elements = append(s.elements, el)
}

func (s *stack) pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("empty stack")
	}
	el := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return el, nil
}

func (s *stack) peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *stack) isEmpty() bool {
	return len(s.elements) == 0
}

// s: []
// sortedStack: [6 7 10]

func sortMinStack(s stack) stack {
	if s.isEmpty() {
		return s
	}
	sortedStack := NewStack()
	for !s.isEmpty() {
		next, _ := s.pop()
		sortedStack = sortMinStackElement(next, sortedStack)
	}
	return sortedStack
}

// el: 7
// s: [6 7 10]
// next:

// el: 7
// s: [7 10]
// next:

func sortMinStackElement(el int, s stack) stack {
	if s.isEmpty() {
		s.push(el)
		return s
	}
	peek, _ := s.peek()
	if el > peek {
		next, _ := s.pop()
		s = sortMinStackElement(el, s)
		s.push(next)
	} else {
		s.push(el)
	}

	return s
}
