package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSumReverseLinkedLists struct {
	name     string
	inList1  *Node
	inList2  *Node
	wantList *Node
}

func TestSumReverseLinkedLists(t *testing.T) {
	// create list123
	list123 := &Node{data: 3, next: &Node{data: 2, next: &Node{data: 1}}}
	// create list246
	list246 := &Node{data: 6, next: &Node{data: 4, next: &Node{data: 2}}}

	// create list19
	list19 := &Node{data: 9, next: &Node{data: 1}}
	// create list2
	list2 := &Node{data: 2}
	// create list21
	list21 := &Node{data: 1, next: &Node{data: 2}}

	// create list879
	list879 := &Node{data: 9, next: &Node{data: 7, next: &Node{data: 8}}}
	// create list586
	list586 := &Node{data: 6, next: &Node{data: 8, next: &Node{data: 5}}}
	// create list1465
	list1465 := &Node{data: 5, next: &Node{data: 6, next: &Node{data: 4, next: &Node{data: 1}}}}

	// create list999
	list999 := &Node{data: 9, next: &Node{data: 9, next: &Node{data: 9}}}
	// create list1001
	list1001 := &Node{data: 1, next: &Node{data: 0, next: &Node{data: 0, next: &Node{data: 1}}}}

	testCases := []testSumReverseLinkedLists{
		{
			name:     "123 + 123 = 246",
			inList1:  list123,
			inList2:  list123,
			wantList: list246,
		},
		{
			name:     "19 + 2 = 21",
			inList1:  list19,
			inList2:  list2,
			wantList: list21,
		},
		{
			name:     "879 + 586 = 1465",
			inList1:  list879,
			inList2:  list586,
			wantList: list1465,
		},
		{
			name:     "999 + 2 = 1001",
			inList1:  list999,
			inList2:  list2,
			wantList: list1001,
		},
	}
	for _, testCase := range testCases {
		// Local version of testCase variable for parallel test execution
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotList := sumReverseLinkedLists(tc.inList1, tc.inList2)
			assert.Equal(t, tc.wantList, gotList)
			// if !reflect.DeepEqual(tc.wantList, gotList) {
			// 	t.Errorf("expected: %v, got: %v", tc.wantList, gotList)
			// }
		})
	}
}
