package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testIsBalanced struct {
	name         string
	inRoot       *node
	wantBalanced bool
}

//                                 n
//                    a                          b
//              c                        d               e
//        f                          g                          h
//                                       i

func TestIsBalanced(t *testing.T) {
	tc1 := &node{left: &node{left: &node{left: &node{}}}, right: &node{left: &node{left: &node{right: &node{}}}, right: &node{right: &node{}}}}
	tc2 := &node{left: &node{left: &node{}}, right: &node{left: &node{}, right: &node{}}}
	tc3 := &node{left: &node{left: &node{}}}
	tc4 := &node{right: &node{left: &node{}}}

	testCases := []testIsBalanced{
		{
			name:         "not balanced, nil input",
			inRoot:       nil,
			wantBalanced: false,
		},
		{
			name:         "not balanced, case 1",
			inRoot:       tc1,
			wantBalanced: false,
		},
		{
			name:         "balanced, case 2",
			inRoot:       tc2,
			wantBalanced: true,
		},
		{
			name:         "not balanced, case 3",
			inRoot:       tc3,
			wantBalanced: false,
		},
		{
			name:         "not balanced, case 4",
			inRoot:       tc4,
			wantBalanced: false,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotBalanced := isBalanced(tc.inRoot)
			assert.Equal(t, tc.wantBalanced, gotBalanced)
		})
	}
}
