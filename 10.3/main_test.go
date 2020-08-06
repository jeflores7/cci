package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSearchRotated struct {
	name    string
	inArr   []int
	inValue int
	wantPos int
}

func TestSearchRotated(t *testing.T) {
	testCases := []testCaseSearchRotated{
		{
			name:    "5 rotations, 8 pos",
			inArr:   []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14},
			inValue: 5,
			wantPos: 8,
		},
		{
			name:    "5 rotations, 2 pos",
			inArr:   []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14},
			inValue: 19,
			wantPos: 2,
		},
		{
			name:    "7 rotations, 10 pos",
			inArr:   []int{10, 14, 15, 16, 19, 20, 25, 1, 3, 4, 5, 7},
			inValue: 5,
			wantPos: 10,
		},
		{
			name:    "12 rotations, 12 pos",
			inArr:   []int{15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14},
			inValue: 1,
			wantPos: 12,
		},
		{
			name:    "12 rotations, 1 pos",
			inArr:   []int{15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14},
			inValue: 16,
			wantPos: 1,
		},
		{
			name:    "12 rotations, 0 pos",
			inArr:   []int{15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14},
			inValue: 15,
			wantPos: 0,
		},
		{
			name:    "12 rotations, 18 pos",
			inArr:   []int{15, 16, 19, 20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14},
			inValue: 14,
			wantPos: 18,
		},
		{
			name:    "2 rotations, 0 pos",
			inArr:   []int{45, 55, 1, 3, 4, 5, 7, 10, 14, 15, 16, 19, 20, 25, 30, 32, 37, 39, 42},
			inValue: 45,
			wantPos: 0,
		},
		{
			name:    "2 rotations, 18 pos",
			inArr:   []int{45, 55, 1, 3, 4, 5, 7, 10, 14, 15, 16, 19, 20, 25, 30, 32, 37, 39, 42},
			inValue: 42,
			wantPos: 18,
		},
		{
			name:    "8 rotations, 5 pos",
			inArr:   []int{20, 25, 30, 32, 37, 39, 42, 45, 55, 1, 3, 4, 5, 7, 10, 14, 15, 16, 19},
			inValue: 39,
			wantPos: 5,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotPos := searchRotated(tc.inArr, tc.inValue)
			assert.Equal(t, tc.wantPos, gotPos)
		})
	}
}
