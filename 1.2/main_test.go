package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		name     string
		a        string
		b        string
		wantBool bool
	}{
		{
			name:     "helloworld, worldhello",
			a:        "helloworld",
			b:        "worldhello",
			wantBool: true,
		},
		{
			name:     "helloworld, helloworld",
			a:        "helloworld",
			b:        "helloworld",
			wantBool: true,
		},
		{
			name:     "helloworld, helloworldplus",
			a:        "helloworld",
			b:        "helloworldplus",
			wantBool: false,
		},
		{
			name:     "helloworldplus, helloworld",
			a:        "helloworldplus",
			b:        "helloworld",
			wantBool: false,
		},
		{
			name:     "helloworld, notworldhello",
			a:        "helloworld",
			b:        "notworldhello",
			wantBool: false,
		},
		{
			name:     "empty first string",
			a:        "",
			b:        "notworldhello",
			wantBool: false,
		},
		{
			name:     "empty second string",
			a:        "abcd",
			b:        "",
			wantBool: false,
		},
		{
			name:     "empty both strings",
			a:        "",
			b:        "",
			wantBool: false,
		},
		{
			name:     "awesome example",
			a:        "aewmeoseexlapm",
			b:        "exampleawesome",
			wantBool: true,
		},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotBool := isPermutation(tc.a, tc.b)
			assert.Equal(t, tc.wantBool, gotBool)
		})
	}
}
