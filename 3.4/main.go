package main

import "errors"

type Stack struct {
	top *Node
}

func (s *Stack) pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("empty stack")
	}
	return s.top.data, nil
}

func (s *Stack) push(data int) {
	newNode := &Node{data: data}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) peek() (int, error) {
	if s.top == nil {
		return 0, errors.New("empty stack")
	}
	return s.top.data, nil
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

type Node struct {
	data int
	next *Node
}

type MyQueue struct {
	data *Stack
	tmp  *Stack
}

func (mq *MyQueue) pop() (int, error) {
	if mq.data == nil || mq.data.isEmpty() {
		return 0, errors.New("empty queue")
	}
	for data, err := mq.data.pop(); err == nil; {
		mq.tmp.push(data)
	}
	data, _ := mq.tmp.pop()
	for data, err := mq.tmp.pop(); err == nil; {
		mq.data.push(data)
	}
	return data, nil
}

func (mq *MyQueue) push(data int) {
	if mq.data == nil {
		mq.data = &Stack{}
	}
	mq.data.push(data)
}

func (mq *MyQueue) peek() (int, error) {
	if mq.data == nil || mq.data.isEmpty() {
		return 0, errors.New("empty queue")
	}
	for data, err := mq.data.pop(); err == nil; {
		mq.tmp.push(data)
	}
	data, _ := mq.tmp.peek()
	for data, err := mq.tmp.pop(); err == nil; {
		mq.data.push(data)
	}
	return data, nil
}

func (mq *MyQueue) isEmpty() bool {
	return mq.data == nil || mq.data.isEmpty()
}
