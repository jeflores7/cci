package main

import (
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
