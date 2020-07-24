package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseIsBST struct {
	name   string
	inRoot *node
	want   bool
}

func TestIsBST(t *testing.T) {
	r1 := &node{
		data: 7,
		left: &node{
			data: 5,
			left: &node{
				data: 2,
			},
		},
		right: &node{
			data: 10,
			left: &node{
				data: 8,
				right: &node{
					data: 9,
				},
			},
			right: &node{
				data: 13,
			},
		},
	}
	r2 := &node{
		data: 8,
		left: &node{
			data: 10,
			right: &node{
				data: 11,
			},
		},
		right: &node{
			data: 9,
		},
	}
	r3 := &node{
		data: 10,
		left: &node{
			data: 5,
			right: &node{
				data: 7,
			},
		},
		right: &node{
			data: 10,
			right: &node{
				data: 12,
			},
		},
	}
	r4 := &node{
		data: 10,
		left: &node{
			data: 7,
			left: &node{
				data: 5,
			},
			right: &node{
				data: 11,
			},
		},
		right: &node{
			data: 14,
			left: &node{
				data: 12,
			},
		},
	}
	r5 := &node{
		data: 10,
		left: &node{
			data: 7,
			right: &node{
				data: 8,
			},
		},
		right: &node{
			data: 15,
			left: &node{
				data: 12,
				left: &node{
					data: 10,
				},
			},
			right: &node{
				data: 17,
			},
		},
	}

	testCases := []testCaseIsBST{
		{
			name:   "r1: yes",
			inRoot: r1,
			want:   true,
		},
		{
			name:   "nil in: yes",
			inRoot: nil,
			want:   true,
		},
		{
			name:   "r2: left child fail: no",
			inRoot: r2,
			want:   false,
		},
		{
			name:   "r3: right child fail: no",
			inRoot: r3,
			want:   false,
		},
		{
			name:   "r4: left indirect child fail: no",
			inRoot: r4,
			want:   false,
		},
		{
			name:   "r5: right indirect child fail: no",
			inRoot: r5,
			want:   false,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := isBST(tc.inRoot)
			assert.Equal(t, tc.want, got)
		})
	}
}
