package main

// Bit operations

// Returns true if the bit at position i is 1, otherwise false
func getBit(num, i int) bool {
	return (num & (1 << i)) != 0
}

// Set the bit at position i to 1
func setBit(num, i int) int {
	return num | (1 << i)
}

// Set the bit at position i to 0
func clearBit(num, i int) int {
	return num & ^(1 << i)
}

// If toOne is true, set the bit at position i to 1, otherwise set the bit at position i to 0
func updateBit(num, i int, toOne bool) int {
	b := 0
	if toOne {
		b = 1
	}
	return num & ^(1<<i) | (b << i)
}

// Flip the bit at position i (0 becomes 1, 1 becomes 0)
func flipBit(num, i int) int {
	return num ^ (1 << i)
}
