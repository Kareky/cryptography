// Package alphabet provides a set of helper functions to work with the alphabet,
// such as converting letters to their positions and vice versa.
package alphabet

// ShiftLetter shifts a letter by the given shift within the alphabet.
// The isUpper flag indicates whether the original letter is uppercase.
// It returns the shifted rune.
func ShiftLetter(letter rune, shift int, isUpper bool) rune {
	initialPosition := FromAlphabetToPosition(letter, isUpper)
	finalPosition := ((initialPosition - 1 + shift) % 26 + 26) % 26 + 1
	shiftedLetter := FromPositionToAlphabet(finalPosition, isUpper)
	return shiftedLetter
}

// FromAlphabetToPosition returns the 1-based position of a letter in the alphabet.
// The isUpper flag indicates whether the letter is uppercase.
func FromAlphabetToPosition(letter rune, isUpper bool) int {
	if !isUpper {
		return int(letter - 'a') + 1
	}
	
	return int(letter - 'A') + 1
}

// FromPositionToAlphabet returns the letter corresponding to the given
// 1-based position in the alphabet. The toUpper flag controls the case of the output.
func FromPositionToAlphabet(position int, toUpper bool) rune {
	if !toUpper {
		return rune('a' + position - 1)
	}
	
	return rune('A' + position - 1)
}