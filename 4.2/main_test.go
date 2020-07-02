package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseMinHeightBST struct {
	name string
	inSlice []int
	wantNode *node
}

func TestMinHeightBST(t *testing.T) {
	testCases := []testCaseMinHeightBST{
		{
			name: "0,1,2,3,4,5",
			inSlice: []int{0,1,2,3,4,5},
			wantNode: &node{data:3,left:&node{data:1,left:&node{data:0},right:&node{data:2}},right:&node{data:5,left:&node{data:4}}},
		},
		{
			name: "nil input",
			inSlice: nil,
			wantNode: nil,
		},
		{
			name: "empty slice",
			inSlice: []int{},
			wantNode: nil,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			
			gotNode := minHeightBST(tc.inSlice)
			assert.Equal(t, tc.wantNode, gotNode)
		})
	}
}