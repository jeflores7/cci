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

// Questions asked:
// 1. What should the return be?
//    Just bool? Just the permutation? Both?
//    Assuming just bool
// 2. Are both inputs guaranteed to be len() > 0?
// Incorrect assumption/understanding:
// 1. Assumed a partial permutation would be satisfactory
// 2.

// Incomplete / started over due to misunderstanding the question
// Should be testing whether the strings are permutations of one another
// not one being a sub-permutation of the other
func isPermutationIncomplete(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	if s1 == s2 {
		return true
	}
	smaller := s2
	bigger := s1
	if len(s2) > len(s1) {
		smaller = s1
		bigger = s2
	}
	smaller = sort(smaller)
	bigger = sort(bigger)
	lastBig := bigger[0]
	for i := range smaller {
		// Because assuming it can be a sub-permutation, it's okay to have more
		// repeated characters in bigger string
		if smaller[i] != bigger[i] && lastBig != bigger[i] {
			return false
		}
		lastBig = bigger[i]
	}
	return true
}

func sort(s string) string {
	// TODO implement
	return ""
}

// Missed:
// 1. Should add a comparison ensuring strings are the same length
//   * This would make it so there is no need to check the hash table for any remaining positive values
func isPermutation(a, b string) bool {
	if len(a) == 0 || len(b) == 0 {
		// Should ask whether empty string is a permutation of empty string?
		return false
	}
	if a == b {
		return true
	}
	aHT := makeHashTable(a)
	for _, r := range b {
		idx := hash(r) % len(aHT)
		if v := aHT[idx]; v-1 < 0 {
			return false
		}
		aHT[idx]--
	}
	// Missed on paper
	for _, v := range aHT {
		if v > 0 {
			return false
		}
	}
	return true
}

func hash(r rune) int {
	// Not implemented on paper
	// Simplified example, assuming ascii lower case alphabet chars only
	// Need to learn more about hashing functions
	return int(r - 'a')
}

// Note:
// 1. Instead of making the hashtable len(a), if we know the alphabet
//    it would be less collisions if we used len(alphabet) instead
func makeHashTable(a string) []int {
	// Not implemented on paper
	ht := make([]int, len(a))
	for _, r := range a {
		idx := hash(r) % len(ht)
		ht[idx]++
	}
	return ht
}
