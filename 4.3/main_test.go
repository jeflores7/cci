package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseGetDepthLists struct {
	name    string
	inRoot  *node
	wantMap map[int]*llNode
}

func TestGetDepthLists(t *testing.T) {
	depth3Example := &node{
		data: 10,
		left: &node{
			data: 1,
			left: &node{
				data:  0,
				right: &node{data: -2},
			},
			right: &node{
				data:  2,
				right: &node{data: 12},
			},
		},
		right: &node{
			data: 5,
			left: &node{
				data:  6,
				right: &node{data: 15},
			},
		},
	}
	depth3ExampleMap := map[int]*llNode{
		0: {data: depth3Example},
		1: {data: depth3Example.left, next: &llNode{data: depth3Example.right}},
		2: {data: depth3Example.left.left, next: &llNode{data: depth3Example.left.right, next: &llNode{data: depth3Example.right.left}}},
		3: {data: depth3Example.left.left.right, next: &llNode{data: depth3Example.left.right.right, next: &llNode{data: depth3Example.right.left.right}}},
	}

	testCases := []testCaseGetDepthLists{
		{
			name:    "depth 3 example",
			inRoot:  depth3Example,
			wantMap: depth3ExampleMap,
		},
		{
			name:    "nil root",
			inRoot:  nil,
			wantMap: make(map[int]*llNode),
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotMap := getDepthLists(tc.inRoot)
			assert.Equal(t, tc.wantMap, gotMap)
		})
	}
}
