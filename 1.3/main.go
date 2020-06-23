package main

// Paper note:
// 1. The slice will be updated, so no need to really return the string
// Missed:
// 1. s = shiftChars(s, i+1, n) was first implementation, however, it does not properly
//    shift the characters. It doesn't account for the spaces that were changed already.
// 2. For loop condition was i < n which does not account for the new length added.
// Brute force way, loop through string from position 0 to truelength
// shift all characters to the right 2 times whenever a space is hit
func urlifyBruteForce(s []rune, n int) string {
	numSpaces := 0
	for i := 0; i < (n + numSpaces*2); i++ {
		if s[i] == ' ' {
			s = shiftChars(s, i+1, (n-1)+numSpaces*2)
			numSpaces++
			s[i] = '%'
			s[i+1] = '2'
			s[i+2] = '0'
		}
	}
	return string(s)
}

func shiftChars(s []rune, start, end int) []rune {
	for i := end; i >= start; i-- {
		s[i+2] = s[i]
	}
	return s
}

// Didn't finish thinking, but on the path
// 1. Counting # of spaces first prevents the need for multiple shifting for each space encountered
//    Instead, every time a space is encountered, the shiftTo position is just adjusted accordingly
// Missed on paper:
// 1. Loop should start at n-1
// 2. Decrement needs to happen for both branches
func urlify(s []rune, n int) {
	var i, numSpaces int
	for i = 0; i < n; i++ {
		if s[i] == ' ' {
			numSpaces++
		}
	}
	shiftTo := (n - 1) + numSpaces*2
	for i = n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s[shiftTo] = '0'
			s[shiftTo-1] = '2'
			s[shiftTo-2] = '%'
			shiftTo -= 2
		} else {
			s[shiftTo] = s[i]
		}
		shiftTo--
	}
	return
}
