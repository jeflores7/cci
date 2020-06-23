package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRemoveDupes struct {
	name            string
	inList          *Node
	inListNoExtraDS *Node
	wantList        *Node
}

func TestRemoveDupes(t *testing.T) {
	ll := &Node{
		Value: 1,
	}
	ll.Insert(1).Insert(2).Insert(3).Insert(4).Insert(4).Insert(5).Insert(6).Insert(7).Insert(7).Insert(8).Insert(9).Insert(10)
	ll2 := &Node{
		Value: 1,
	}
	ll2.Insert(1).Insert(2).Insert(3).Insert(4).Insert(4).Insert(5).Insert(6).Insert(7).Insert(7).Insert(8).Insert(9).Insert(10)
	llNoDupes := &Node{
		Value: 1,
	}
	llNoDupes.Insert(2).Insert(3).Insert(4).Insert(5).Insert(6).Insert(7).Insert(8).Insert(9).Insert(10)

	allDupes := &Node{
		Value: 1,
	}
	allDupes2 := &Node{
		Value: 1,
	}
	last := allDupes
	last2 := allDupes2
	for i := 0; i < 100000; i++ {
		last = last.Insert(1)
		last2 = last2.Insert(1)
	}
	allDupesClean := &Node{
		Value: 1,
	}
	someCollisions := &Node{
		Value: 1,
	}
	someCollisions.Insert(-1).Insert(2).Insert(-2).Insert(2).Insert(-2).Insert(3).Insert(-3).Insert(4).Insert(-4).
		Insert(-4).Insert(-4).Insert(-4).Insert(-5).Insert(-7).Insert(-7).Insert(-9).Insert(10)
	someCollisions2 := &Node{
		Value: 1,
	}
	someCollisions2.Insert(-1).Insert(2).Insert(-2).Insert(2).Insert(-2).Insert(3).Insert(-3).Insert(4).Insert(-4).
		Insert(-4).Insert(-4).Insert(-4).Insert(-5).Insert(-7).Insert(-7).Insert(-9).Insert(10)
	someCollisionsClean := &Node{
		Value: 1,
	}
	someCollisionsClean.Insert(-1).Insert(2).Insert(-2).Insert(3).Insert(-3).Insert(4).Insert(-4).Insert(-5).
		Insert(-7).Insert(-9).Insert(10)

	collisionsFurtherApart := &Node{
		Value: 1,
	}
	collisionsFurtherApart.Insert(2).Insert(3).Insert(-4).Insert(4).Insert(-5).Insert(5).Insert(5).Insert(-5).
		Insert(4).Insert(-4).Insert(3).Insert(2).Insert(1)
	collisionsFurtherApart2 := &Node{
		Value: 1,
	}
	collisionsFurtherApart2.Insert(2).Insert(3).Insert(-4).Insert(4).Insert(-5).Insert(5).Insert(5).Insert(-5).
		Insert(4).Insert(-4).Insert(3).Insert(2).Insert(1)
	collisionsFurtherApartClean := &Node{
		Value: 1,
	}
	collisionsFurtherApartClean.Insert(2).Insert(3).Insert(-4).Insert(4).Insert(-5).Insert(5)
	testCases := []testRemoveDupes{
		{
			name:            "1 1 2 3 4 4 5 6 7 7 8 9 10",
			inList:          ll,
			inListNoExtraDS: ll2,
			wantList:        llNoDupes,
		},
		{
			name:            "empty list",
			inList:          nil,
			inListNoExtraDS: nil,
			wantList:        nil,
		},
		{
			name:            "All dupes 1 1 1 1 1 1 1 1 1 1 ...",
			inList:          allDupes,
			inListNoExtraDS: allDupes2,
			wantList:        allDupesClean,
		},
		{
			name:            "Some collisions: 1 -1 2 -2 2 -2 3 -3 4 -4 -4 -4 -4 -5 -7 -7 -9 10",
			inList:          someCollisions,
			inListNoExtraDS: someCollisions2,
			wantList:        someCollisionsClean,
		},
		{
			name:            "Collisions further apart: 1 2 3 -4 4 -5 5 5 -5 4 -4 3 2 1",
			inList:          collisionsFurtherApart,
			inListNoExtraDS: collisionsFurtherApart2,
			wantList:        collisionsFurtherApartClean,
		},
	}

	for _, testCase := range testCases {
		// Local version of testCase variable for parallel execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			removeDupes(tc.inList)
			assert.Equal(t, tc.wantList, tc.inList)

			removeDupesNoExtraDS(tc.inListNoExtraDS)
			assert.Equal(t, tc.wantList, tc.inListNoExtraDS)
		})
	}
}
