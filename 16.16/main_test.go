package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// in := 1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19
// expect: start = 3, end = 9

// i := 3
// in[i] := 7

// max := 19
// min := 6

// start := 3
// end := 9

// --

// in := 10, 1, 2, 3, 4, 5
// expect: start = 0, end = 5

// i := 0
// in[i] := 10

// max := 10
// min := 1

// start := 0
// end := 5

// --

// in := 0, 1, 2, 3
// expect: start = end = 0, err != nil

// i := 4
// in[i] := -

// max := 3
// min := 3

// start := 0
// end := 0
// err := errors.New("")

type testCaseGetUnsortedIndexes struct {
	name      string
	inArr     []int
	wantStart int
	wantEnd   int
	wantErr   error
}

func TestGetUnsortedIndexes(t *testing.T) {
	testCases := []testCaseGetUnsortedIndexes{
		{
			name:      "random list",
			inArr:     []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19},
			wantStart: 3,
			wantEnd:   9,
			wantErr:   nil,
		},
		{
			name:      "special case: entire array",
			inArr:     []int{10, 1, 2, 3, 4, 5},
			wantStart: 0,
			wantEnd:   5,
			wantErr:   nil,
		},
		{
			name:      "special case: already sorted array",
			inArr:     []int{0, 1, 2, 3, 4, 5},
			wantStart: 0,
			wantEnd:   0,
			wantErr:   errors.New(""),
		},
		{
			name:      "special case: array too short",
			inArr:     []int{1},
			wantStart: 0,
			wantEnd:   0,
			wantErr:   errors.New(""),
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotStart, gotEnd, gotErr := getUnsortedIndexes(tc.inArr)
			assert.Equal(t, tc.wantStart, gotStart)
			assert.Equal(t, tc.wantEnd, gotEnd)
			if tc.wantErr == nil {
				assert.Nil(t, gotErr)
			} else {
				assert.NotNil(t, gotErr)
			}
		})
	}
}
