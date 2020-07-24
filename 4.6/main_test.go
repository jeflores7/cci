package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseFindNext struct {
	name     string
	inNode   *node
	wantNode *node
}

func populateParents(node *node, parent *node) {
	if node == nil {
		return
	}
	node.parent = parent
	populateParents(node.left, node)
	populateParents(node.right, node)
}

func TestFindNext(t *testing.T) {
	rightChildTreeEnd := &node{
		data: 11,
	}
	rightChildTree := &node{
		data: 10,
		right: &node{
			data: 15,
			left: &node{
				data: 13,
				left: rightChildTreeEnd,
				right: &node{
					data: 14,
				},
			},
		},
	}
	populateParents(rightChildTree, nil)

	ancestorStart := &node{
		data: 9,
		left: &node{
			data: 7,
		},
	}
	ancestorTree := &node{
		data: 10,
		left: &node{
			data: 5,
			right: &node{
				data:  7,
				right: ancestorStart,
			},
		},
	}
	populateParents(ancestorTree, nil)

	testCases := []testCaseFindNext{
		{
			name:     "nil input",
			inNode:   nil,
			wantNode: nil,
		},
		{
			name:     "right child",
			inNode:   rightChildTree,
			wantNode: rightChildTreeEnd,
		},
		{
			name:     "ancestor",
			inNode:   ancestorStart,
			wantNode: ancestorTree,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNode := findNext(tc.inNode)
			assert.Equal(t, tc.wantNode, gotNode)
		})
	}
}
