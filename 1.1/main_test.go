package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "lower case ascii, sorted",
			input:          "abcdefghijklmnopqrstuvwxyz",
			expectedResult: true,
		},
		{
			name:           "abcd repeated, sorted",
			input:          "aaabbbcccddd",
			expectedResult: false,
		},
		{
			name:           "a repeated at begin and end, mostly sorted",
			input:          "abcdefghijkla",
			expectedResult: false,
		},
		{
			name:           "a repeated at begin and end, unsorted",
			input:          "awruvonpqejia",
			expectedResult: false,
		},
		{
			name:           "z repeated at begin and end, unsorted",
			input:          "zwruvonpqejiz",
			expectedResult: false,
		},
		{
			name:           "all lower case alphabet characters, some repeats, unsorted",
			input:          "thequickbrownfoxjumpsoverthelazydog",
			expectedResult: false,
		},
		{
			name:           "all lower case alphabet characters, no repeats, unsorted",
			input:          "cwmfjordbankglyphsvextquiz",
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		// Local instance of testCase for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			isUnique := uniqueChars(tc.input)
			assert.Equal(t, isUnique, tc.expectedResult)

			isUnique = uniqueCharsNoExtraDS(tc.input)
			assert.Equal(t, isUnique, tc.expectedResult)
		})
	}
}

// Questions should've asked:
// 1. How big is the alphabet?
//   - Can compare length of string to alphabet size and return early if it larger due to repeated chars
// 2. Learn about bit vectors and how to use them
// Changes could've made:
// 1. Instead of a map, just could've used a []int the size of the alphabet
//    Then, each character maps to an index of the slice and if any value grows greater than 1
// 2. Same as above, but just use bool instead
func uniqueChars(s string) bool {
	charsSeen := make(map[rune]struct{}, 0)

	for _, r := range s {
		if _, found := charsSeen[r]; found {
			// If we've seen a character already, we do not have unique characters
			return false
		}
		charsSeen[r] = struct{}{}
	}

	return true
}

// Slower, but no extra structures
// Notes
// 1. Had to lookup sort.Slice, get more familiar with sort package
// 2. Had index out of range error due to reaching end of rune slice and comparing with i+1
func uniqueCharsNoExtraDS(s string) bool {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	numRunes := len(runes)
	for i, r := range runes {
		if i < numRunes-1 && r == runes[i+1] {
			return false
		}
	}
	return true
}
