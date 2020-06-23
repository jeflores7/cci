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
