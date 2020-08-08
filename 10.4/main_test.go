package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSearchListy struct {
	name      string
	inListy   list
	inValue   int
	wantValue int
}

func TestSearchListy(t *testing.T) {
	max16 := make([]int, 0, 17)
	for i := 0; i <= 16; i++ {
		max16 = append(max16, i)
	}
	max32 := make([]int, 0, 33)
	for i := 0; i <= 32; i++ {
		max32 = append(max32, i)
	}
	testCases := []testCaseSearchListy{
		{
			name: "search less than 0",
			inListy: list{
				data: max16,
			},
			inValue:   -1,
			wantValue: -1,
		},
		{
			name: "search 0",
			inListy: list{
				data: max16,
			},
			inValue:   0,
			wantValue: 0,
		},
		{
			name: "max: 2^4, find: 2^3 + 5",
			inListy: list{
				data: max16,
			},
			inValue:   13,
			wantValue: 13,
		},
		{
			name: "max: 2^5, find: 2^4 + 10",
			inListy: list{
				data: max32,
			},
			inValue:   56,
			wantValue: -1,
		},
		{
			name: "max: 2^5, find: >2^5",
			inListy: list{
				data: max32,
			},
			inValue:   72,
			wantValue: -1,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotValue := searchListy(tc.inListy, tc.inValue)
			assert.Equal(t, tc.wantValue, gotValue)
		})
	}
}
