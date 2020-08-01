package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCountWays struct {
	name  string
	inN   int
	wantN int
}

func TestCountWays(t *testing.T) {
	testCases := []testCountWays{
		{
			name:  "n = 12",
			inN:   12,
			wantN: 4,
		},
		{
			name:  "n = 27",
			inN:   27,
			wantN: 13,
		},
		{
			name:  "n = 100",
			inN:   100,
			wantN: 213,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotN := countWays(tc.inN)
			assert.Equal(t, tc.wantN, gotN)
		})
	}
}
