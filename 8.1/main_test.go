package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTripleStep struct {
	name  string
	inN   int
	wantN int
}

func TestTripleStep(t *testing.T) {
	testCases := []testTripleStep{
		{
			name:  "3: 4",
			inN:   3,
			wantN: 4,
		},
		{
			name:  "4: 7",
			inN:   4,
			wantN: 7,
		},
		{
			name:  "0: 0",
			inN:   0,
			wantN: 0,
		},
		{
			name:  "10: 38",
			inN:   10,
			wantN: 274,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotN := tripleStep(tc.inN)
			assert.Equal(t, tc.wantN, gotN)
		})
	}
}
