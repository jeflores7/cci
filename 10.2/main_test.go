package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseAnagramSort struct {
	name      string
	inSlice   []string
	wantSlice []string
}

func TestAnagramSort(t *testing.T) {
	testCases := []testCaseAnagramSort{
		{
			name:      "sample",
			inSlice:   []string{"hello", "world", "asdfh", "ehllo", "beicn", "drolw"},
			wantSlice: []string{"hello", "ehllo", "world", "drolw", "asdfh", "beicn"},
		},
		{
			name:      "double sample",
			inSlice:   []string{"hello", "world", "asdfh", "ehllo", "beicn", "drolw", "hello", "world", "asdfh", "ehllo", "beicn", "drolw"},
			wantSlice: []string{"hello", "ehllo", "hello", "ehllo", "world", "drolw", "world", "drolw", "asdfh", "asdfh", "beicn", "beicn"},
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			anagramSort(tc.inSlice)
			assert.Equal(t, tc.wantSlice, tc.inSlice)
		})
	}
}
