package main

import "strings"

// Brute force
// For each, compare to every other string
// if anagram, append to list of anagrams
// Append all lists
// Use hash table
// Sort each string to get key in hash table
// Append string to list of strings at that key
// Append all lists
// Given array of strings with length n
// Given all strings are of len s
// Given all strings have ascii characters of english alphabet [a-zA-z]
// Given case insensitive anagrams
// O(sn)

// s: [ "hello" "world" "elloh" "drolw" "asdfh" beicn" ]
// sKeys: {

// }
// i: 0

func anagramSort(s []string) {
	sKeys := make(map[string][]string)
	keysOrder := make([]string, 0) // For stable testing
	for i := range s {
		key := countSort(s[i])
		if _, found := sKeys[key]; !found {
			keysOrder = append(keysOrder, key)
		}
		sKeys[key] = append(sKeys[key], s[i])
	}
	i := 0
	for _, key := range keysOrder {
		for j := range sKeys[key] {
			s[i] = sKeys[key][j]
			i++
		}
	}
}

// s: "hello"
// counts: [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// i: 0

func countSort(s string) string {
	max := 26 // 26 characters in alphabet
	counts := make([]int, max)

	for _, i := range s {
		counts[countKey(i)]++
	}

	builder := strings.Builder{}
	builder.Grow(len(s))
	for i := range counts {
		for j := 0; j < counts[i]; j++ {
			builder.WriteRune(asciiHelper(i))
		}
	}

	return builder.String()
}

const letterA int = 97

func countKey(letter rune) int {
	return int(letter) - letterA
}

func asciiHelper(letter int) rune {
	return rune(letterA + letter)
}
