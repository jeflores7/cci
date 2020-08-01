package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseMagicIndex struct {
	name    string
	inArr   []int
	wantIdx int
}

func TestMagicIndex(t *testing.T) {
	testCases := []testCaseMagicIndex{
		{
			name:    "3: [ -1 0 1 3 5 8 11 ]",
			inArr:   []int{-1, 0, 1, 3, 5, 8, 11},
			wantIdx: 3,
		},
		{
			name:    "6: [ -1 0 1 2 3 4 6 ]",
			inArr:   []int{-1, 0, 1, 2, 3, 4, 6},
			wantIdx: 6,
		},
		{
			name:    "-1: [ -1 0 1 2 3 4 7 ]",
			inArr:   []int{-1, 0, 1, 2, 3, 4, 7},
			wantIdx: -1,
		},
		{
			name:    "-1: [ ]",
			inArr:   []int{},
			wantIdx: -1,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotIdx := magicIndex(tc.inArr)
			assert.Equal(t, tc.wantIdx, gotIdx)
		})
	}
}
