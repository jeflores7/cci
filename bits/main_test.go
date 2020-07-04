package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseGetBit struct {
	name     string
	inNum    int
	inBitPos int
	wantBit  bool
}

func TestGetBit(t *testing.T) {
	testCases := []testCaseGetBit{
		{
			name:     "get bit at 4's position from 4 (0100)",
			inNum:    4,
			inBitPos: 2,
			wantBit:  true,
		},
		{
			name:     "get bit at 4's position from -4 (111...0100)",
			inNum:    -4,
			inBitPos: 2,
			wantBit:  true,
		},
		{
			name:     "get bit at 128's position from 129 (10000001)",
			inNum:    129,
			inBitPos: 7,
			wantBit:  true,
		},
		{
			name:     "get bit at 128's position from 260 (100000100)",
			inNum:    260,
			inBitPos: 7,
			wantBit:  false,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNum := getBit(tc.inNum, tc.inBitPos)
			assert.Equal(t, tc.wantBit, gotNum)
		})
	}
}

type testCaseSetBit struct {
	name     string
	inNum    int
	inBitPos int
	wantNum  int
}

func TestSetBit(t *testing.T) {
	testCases := []testCaseSetBit{
		{
			name:     "4 to 260",
			inNum:    4,
			inBitPos: 8,
			wantNum:  260,
		},
		{
			name:     "15 to 79",
			inNum:    15,
			inBitPos: 6,
			wantNum:  79,
		},
		{
			name:     "-4 to -2",
			inNum:    -4,
			inBitPos: 1,
			wantNum:  -2,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNum := setBit(tc.inNum, tc.inBitPos)
			assert.Equal(t, tc.wantNum, gotNum)
		})
	}
}

type testCaseClearBit struct {
	name     string
	inNum    int
	inBitPos int
	wantNum  int
}

func TestClearBit(t *testing.T) {
	testCases := []testCaseClearBit{
		{
			name:     "15 to 11",
			inNum:    15,
			inBitPos: 2,
			wantNum:  11,
		},
		{
			name:     "255 to 191",
			inNum:    255,
			inBitPos: 6,
			wantNum:  191,
		},
		{
			name:     "-4 to -68",
			inNum:    -4,
			inBitPos: 6,
			wantNum:  -68,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNum := clearBit(tc.inNum, tc.inBitPos)
			assert.Equal(t, tc.wantNum, gotNum)
		})
	}
}

type testCaseUpdateBit struct {
	name     string
	inNum    int
	inBitPos int
	inToOne  bool
	wantNum  int
}

func TestUpdateBit(t *testing.T) {
	testCases := []testCaseUpdateBit{
		{
			name:     "4 to 0",
			inNum:    4,
			inBitPos: 2,
			inToOne:  false,
			wantNum:  0,
		},
		{
			name:     "31 to 63",
			inNum:    31,
			inBitPos: 5,
			inToOne:  true,
			wantNum:  63,
		},
		{
			name:     "76 to 108",
			inNum:    76,
			inBitPos: 5,
			inToOne:  true,
			wantNum:  108,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNum := updateBit(tc.inNum, tc.inBitPos, tc.inToOne)
			assert.Equal(t, tc.wantNum, gotNum)
		})
	}
}

type testCaseFlipBit struct {
	name     string
	inNum    int
	inBitPos int
	wantNum  int
}

func TestFlipBit(t *testing.T) {
	testCases := []testCaseFlipBit{
		{
			name:     "4 to 0",
			inNum:    4,
			inBitPos: 2,
			wantNum:  0,
		},
		{
			name:     "108 to 76",
			inNum:    108,
			inBitPos: 5,
			wantNum:  76,
		},
		{
			name:     "15616 to 13568",
			inNum:    15616,
			inBitPos: 11,
			wantNum:  13568,
		},
	}
	for _, testCase := range testCases {
		// Local version of the testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotNum := flipBit(tc.inNum, tc.inBitPos)
			assert.Equal(t, tc.wantNum, gotNum)
		})
	}
}
