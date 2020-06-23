package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testIsPalindrome struct {
	name         string
	inList       *Node
	inListLength int
	want         bool
}

func TestIsPalindrome(t *testing.T) {
	racecar := &Node{data: 'r'}
	racecar.Insert('a').Insert('c').Insert('e').Insert('c').Insert('a').Insert('r')

	raceca := &Node{data: 'r'}
	raceca.Insert('a').Insert('c').Insert('e').Insert('c').Insert('a')

	nope := &Node{data: 'n'}
	nope.Insert('o').Insert('p').Insert('e')

	anutforajaroftuna := &Node{data: 'A'}
	anutforajaroftuna.Insert('n').Insert('u').Insert('t').
		Insert('f').Insert('o').Insert('r').
		Insert('a').
		Insert('j').Insert('a').Insert('r').
		Insert('o').Insert('f').
		Insert('t').Insert('u').Insert('n').Insert('a')

	testCases := []testIsPalindrome{
		{
			name:         "empty list",
			inList:       nil,
			inListLength: 0,
			want:         false,
		},
		{
			name:         "racecar",
			inList:       racecar,
			inListLength: 7,
			want:         true,
		},
		{
			name:         "raceca",
			inList:       raceca,
			inListLength: 6,
			want:         false,
		},
		{
			name:         "nope",
			inList:       nope,
			inListLength: 4,
			want:         false,
		},
		{
			name:         "Anutforajaroftuna",
			inList:       anutforajaroftuna,
			inListLength: 17,
			want:         true,
		},
	}

	for _, testCase := range testCases {
		// Local version of testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := isPalindrome(tc.inList)
			assert.Equal(t, tc.want, got)

			_, got = isPalindromeRecursive(tc.inList, 0, tc.inListLength)
			assert.Equal(t, tc.want, got)
		})
	}
}
