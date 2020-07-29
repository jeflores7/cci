package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFindBuildPath struct {
	name       string
	inProjects []string
	inDeps     [][]string
	wantPath   []string
	wantErr    error
}

func TestFindBuildPath(t *testing.T) {
	testCases := []testFindBuildPath{
		{
			name:       "success, nil",
			inProjects: nil,
			inDeps:     nil,
			wantPath:   nil,
			wantErr:    errors.New(""),
		},
		{
			name:       "success",
			inProjects: []string{"a", "b", "c", "d", "e", "f"},
			inDeps: [][]string{
				{"a", "d"},
				{"f", "b"},
				{"b", "d"},
				{"f", "a"},
				{"d", "c"},
			},
			wantPath: []string{"e", "f", "a", "b", "d", "c"},
			wantErr:  nil,
		},
		{
			name:       "success harder",
			inProjects: []string{"a", "b", "c", "d", "e", "f", "g"},
			inDeps: [][]string{
				{"f", "a"},
				{"f", "b"},
				{"f", "c"},
				{"d", "g"},
				{"c", "a"},
				{"b", "a"},
				{"b", "e"},
				{"a", "e"},
			},
			wantPath: []string{"d", "f", "b", "c", "a", "e", "g"},
			wantErr:  nil,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotPath, gotErr := findBuildPath(tc.inProjects, tc.inDeps)

			if tc.wantErr == nil {
				assert.Nil(t, gotErr)
			} else {
				assert.NotNil(t, gotErr)
			}
			assert.Equal(t, tc.wantPath, gotPath)
		})
	}
}
