package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		name         string
		inRunes      []rune
		inTrueLength int
		wantURL      string
	}{
		{
			name:         "Mr John Smith    ",
			inRunes:      []rune("Mr John Smith    "),
			inTrueLength: 13,
			wantURL:      "Mr%20John%20Smith",
		},
		{
			name:         "  Mr  J  Smith            ",
			inRunes:      []rune("  Mr  J  Smith            "),
			inTrueLength: 14,
			wantURL:      "%20%20Mr%20%20J%20%20Smith",
		},
		{
			name:         "            a                        ",
			inRunes:      []rune("            a                        "),
			inTrueLength: 13,
			wantURL:      "%20%20%20%20%20%20%20%20%20%20%20%20a",
		},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bfRunes := make([]rune, len(tc.inRunes))
			copy(bfRunes, tc.inRunes)
			gotURL := urlifyBruteForce(bfRunes, tc.inTrueLength)
			assert.Equal(t, tc.wantURL, gotURL)

			urlify(tc.inRunes, tc.inTrueLength)
			assert.Equal(t, tc.wantURL, string(tc.inRunes))
		})
	}
}

// Paper note:
// 1. The slice will be updated, so no need to really return the string
// Missed:
// 1. s = shiftChars(s, i+1, n) was first implementation, however, it does not properly
//    shift the characters. It doesn't account for the spaces that were changed already.
// 2. For loop condition was i < n which does not account for the new length added.
// Brute force way, loop through string from position 0 to truelength
// shift all characters to the right 2 times whenever a space is hit
func urlifyBruteForce(s []rune, n int) string {
	numSpaces := 0
	for i := 0; i < (n + numSpaces*2); i++ {
		if s[i] == ' ' {
			s = shiftChars(s, i+1, (n-1)+numSpaces*2)
			numSpaces++
			s[i] = '%'
			s[i+1] = '2'
			s[i+2] = '0'
		}
	}
	return string(s)
}

func shiftChars(s []rune, start, end int) []rune {
	for i := end; i >= start; i-- {
		s[i+2] = s[i]
	}
	return s
}

// Didn't finish thinking, but on the path
// 1. Counting # of spaces first prevents the need for multiple shifting for each space encountered
//    Instead, every time a space is encountered, the shiftTo position is just adjusted accordingly
// Missed on paper:
// 1. Loop should start at n-1
// 2. Decrement needs to happen for both branches
func urlify(s []rune, n int) {
	var i, numSpaces int
	for i = 0; i < n; i++ {
		if s[i] == ' ' {
			numSpaces++
		}
	}
	shiftTo := (n - 1) + numSpaces*2
	for i = n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s[shiftTo] = '0'
			s[shiftTo-1] = '2'
			s[shiftTo-2] = '%'
			shiftTo -= 2
		} else {
			s[shiftTo] = s[i]
		}
		shiftTo--
	}
	return
}
