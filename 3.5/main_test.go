package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSortMinStack struct {
	name      string
	inStack   stack
	wantStack stack
}

func TestSortMinStack(t *testing.T) {
	t1 := NewStack()
	t1.push(7)
	t1.push(6)
	t1.push(10)
	s1 := NewStack()
	s1.push(10)
	s1.push(7)
	s1.push(6)

	t2 := NewStack()
	t2.push(9)
	t2.push(3)
	t2.push(8)
	t2.push(4)
	t2.push(1)
	t2.push(2)
	t2.push(5)
	t2.push(7)
	t2.push(6)
	t2.push(10)
	s2 := NewStack()
	s2.push(10)
	s2.push(9)
	s2.push(8)
	s2.push(7)
	s2.push(6)
	s2.push(5)
	s2.push(4)
	s2.push(3)
	s2.push(2)
	s2.push(1)
	testCases := []testCaseSortMinStack{
		{
			name:      "10 6 7 -> 6 7 10",
			inStack:   t1,
			wantStack: s1,
		},
		{
			name:      "empty input",
			inStack:   NewStack(),
			wantStack: NewStack(),
		},
		{
			name:      "10 6 7 5 2 1 4 8 3 9 -> 1 2 3 4 5 6 7 8 9 10",
			inStack:   t2,
			wantStack: s2,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotStack := sortMinStack(tc.inStack)
			assert.Equal(t, tc.wantStack, gotStack)
		})
	}
}
