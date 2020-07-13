package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseLongestValid struct {
	name    string
	inStr   string
	wantStr string
}

func TestLongestValid(t *testing.T) {
	testCases := []testCaseLongestValid{
		{
			name:    "bxoxok: book",
			inStr:   "bxoxok",
			wantStr: "book",
		},
		{
			name:    "haaoaaraaiaazaaoaan: horizon",
			inStr:   "haaoaaraaiaazaaoaan",
			wantStr: "horizon",
		},
		{
			name:    "heroalwayswins: always",
			inStr:   "heroalwayswins",
			wantStr: "always",
		},
		{
			name:    "haerolwianys: always",
			inStr:   "haerolwianys",
			wantStr: "always",
		},
		{
			name:    "supercalifragilisticexpialidocious: supercalifragilisticexpialidocious",
			inStr:   "supercalifragilisticexpialidocious",
			wantStr: "supercalifragilisticexpialidocious",
		},
		{
			name:    "estanbtlnipsmhzmpecnltv: establishment",
			inStr:   "estanbtlnipsmhzmpecnltv",
			wantStr: "establishment",
		},
		{
			name:    "qrgqrtqbfgqwrrtnwmxm: \"\"",
			inStr:   "qrgqrtqbfgqwrrtnwmxm",
			wantStr: "",
		},
		{
			name:    "aaaaaaaaaaaaaaaaaaaaaaa: ",
			inStr:   "aaaaaaaaaaaaaaaaaaaaaaa",
			wantStr: "a",
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotStr := LongestValid(tc.inStr)
			assert.Equal(t, tc.wantStr, gotStr)
		})
	}
}
