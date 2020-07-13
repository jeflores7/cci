package main

// IsValid(string) -> bool

// IsValid("book") -> true
// IsValid("bbook") -> false

// LongestValid(string) -> string
// LongestValid("bxoxok") -> "book"

var validHardCoded = map[string]struct{}{
	"":                                   {},
	"a":                                  {},
	"always":                             {},
	"book":                               {},
	"callous":                            {},
	"ear":                                {},
	"establishment":                      {},
	"fragile":                            {},
	"hero":                               {},
	"hear":                               {},
	"hello":                              {},
	"heroism":                            {},
	"horizon":                            {},
	"i":                                  {},
	"super":                              {},
	"supercalifragilisticexpialidocious": {},
	"wins":                               {},
}

// IsValid Awesome function that can determine if the given string is a valid string
func IsValid(in string) bool {
	if _, found := validHardCoded[in]; found {
		return true
	}
	return false
}

// in: bxoxok
// i := 0, substr :=  + xoxok
// longest := book

// in: boxok
// i := 2, substr := bo + ok

// in: oxok
// in: xok
// in: ok

// next = [ oxok, xxok, xook, xoxk, xoxo, oxok, bxok, book, boxk, boxo ]
// i = 1
// cur = bxoxok
// substr = b + oxok

// LongestValid Takes any string and finds the longest valid string within it
func LongestValid(in string) string {
	if IsValid(in) {
		return in
	}
	testedStrings := make(map[string]struct{})
	testedStrings[in] = struct{}{}
	next := make([]string, 1)
	next[0] = in
	for len(next) > 0 {
		cur := next[0]
		next = next[1:]
		if IsValid(cur) {
			return cur
		}
		if len(cur) == 1 {
			continue
		}
		for i := 0; i < len(cur); i++ {
			substr := ""
			if i == 0 {
				substr = cur[1:]
			} else {
				substr = cur[0:i] + cur[i+1:]
			}
			if _, found := testedStrings[substr]; !found {
				next = append(next, substr)
				testedStrings[substr] = struct{}{}
			}
			// if IsValid(substr) {
			// 	return substr
			// }
			// next = append(next, getSubstringPermutations(substr, testedStrings)...)
		}
	}
	return ""
}

// func getSubstringPermutations(in string, testedStrings map[string]struct{}) []string {
// 	perms := make([]string, 0, len(in))
// 	for i := 0; i < len(in); i++ {
// 		substr := ""
// 		if i == 0 {
// 			substr = in[1:]
// 		} else {
// 			substr = in[0:i] + in[i+1:]
// 		}
// 		if _, found := testedStrings[substr]; !found {
// 			perms = append(perms, substr)
// 			testedStrings[substr] = struct{}{}
// 		}
// 	}
// 	return perms
// }

// // LongestValid Takes any string and finds the longest valid string within it
// func LongestValid(in string) string {
// 	if IsValid(in) {
// 		return in
// 	}

// 	return longestRecursive(in, "")
// }

// // Check for valid strings by removing one character at a time at this level
// // If none found, recursively call function generating substrings
// func longestRecursive(in string, curLongest string) string {
// 	if IsValid(in) {
// 		return in
// 	}
// 	for i := 0; i < len(in); i++ {
// 		substr := ""
// 		if i == 0 {
// 			substr = in[1:]
// 		} else {
// 			substr = in[0:i] + in[i+1:]
// 		}
// 		if IsValid(substr) {
// 			return substr
// 		}
// 	}
// 	if len(curLongest) > len(in)-1 {
// 		return curLongest
// 	}
// 	longest := ""
// 	for i := 0; i < len(in); i++ {
// 		substr := ""
// 		if i == 0 {
// 			substr = in[1:]
// 		} else {
// 			substr = in[0:i] + in[i+1:]
// 		}
// 		subLongest := longestRecursive(substr, curLongest)
// 		if len(longest) < len(subLongest) {
// 			longest = subLongest
// 		}
// 	}
// 	return longest
// }

// IsValid(string) -> bool

// IsValid(“book”) -> true
// IsValid(“bbook”) -> false

// LongestValid(string) -> string
// LongestValid(“bxoxok”) -> “book”

// input length not too long
// only removing characters, preserving order
// empty string will be valid, if input has no valid strings, return empty string

// Brute force
// Build out all permutations, removing all substrings

// "bxoxok"

// "xoxok"
// "oxok"
// "xok"
// ..

// boxok
// bxok
// bok
// ..

// func LongestValid(in string) string {
// 	if IsValid(in) {
// 		return in
// 	}

// 	longest := buildLongestValid(in, 1)
// 	return longest
// }

// // bxoxok 1

// // oxok 2

// func buildLongestValid(in string, subStringRemovals int) string {
// 	if subStringRemovals < len(in) {
// 		return ""
// 	}
// 	longestValid := ""
// 	for i := 0; i < len(in); i++ {
// 		var sub string
// 		if i == 0 {
// 		sub = string(in[:subStringRemovals])
// 	} else {
// 		sub = string(in[0:i]) + string(in[i+subStringRemovals:])
// 	}
// 		if IsValid(sub) {
// 			longestValid = sub
// 			return longestValid
// 		}
// 	}

// 	return longestValid
// }
