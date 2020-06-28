package main

import "errors"

// Find m & n indices in an array such that sorting m through n would result in a fully sorted array.

// Solutions:
// Sort, then find m/n by comparing to sorted array: O(n log n)
// Single iteration: O(n)

// Steps single iteration:
// Find unsorted elements and their positions
// Index of last unsorted element equals n
// Find expected position of min of unsorted elements
// Equals m

// in := 1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19
// expect: start = 3, end = 9

// i := 3
// in[i] := 7

// max := 19
// min := 6

// start := 3
// end := 9

// --

// in := 10, 1, 2, 3, 4, 5
// expect: start = 0, end = 5

// i := 0
// in[i] := 10

// max := 10
// min := 1

// start := 0
// end := 5

// --

// in := 0, 1, 2, 3
// expect: start = end = 0, err != nil

// i := 4
// in[i] := -

// max := 3
// min := 3

// start := 0
// end := 0
// err := errors.New("")

func getUnsortedIndexes(in []int) (start int, end int, err error) {
	if len(in) < 2 {
		err := errors.New("minimum length not met")
		return start, end, err
	}
	max := in[0]
	min := in[len(in)-1]
	for i := 1; i < len(in); i++ {
		if in[i] < max {
			end = i
			if min > in[i] {
				min = in[i]
			}
			continue
		}
		max = in[i]
	}
	if max == min {
		err := errors.New("array already sorted")
		return start, end, err
	}
	for i := 0; i < len(in); i++ {
		if in[i] > min {
			start = i
			break
		}
	}
	return start, end, nil
}
