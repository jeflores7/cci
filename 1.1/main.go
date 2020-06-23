package main

import "sort"

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
