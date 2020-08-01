package main

import "testing"

type testCasePrintNQueens struct {
	name string
	inN  int
}

func TestPrintNQueens(t *testing.T) {
	testCases := []testCasePrintNQueens{
		{
			name: "8x8",
			inN:  8,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			printNQueens(tc.inN)
		})
	}
}
