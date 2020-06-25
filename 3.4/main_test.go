package main

import "testing"

type testMyQueue struct {
	name string
}

func TestMyQueue(t *testing.T) {
	testCases := []testMyQueue{
		{
			name: "",
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}
