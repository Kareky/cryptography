package caesarCypher

import (
	"strings"
	"unicode"
)

// Encrypt takes a message and encrypts it, moving each character inside the sentence
// according to shift. Uppercase can be preserved or skipped by using preserveUppercase.
// The function returns the shifted string, preserving each space or symbol
// as it was in the original message.
func Encrypt(text string, shift int, preserveUppercase bool) string {
	var newText strings.Builder
	if !preserveUppercase {
		for _, character := range text {
			character = unicode.ToLower(character)
			if character >= 'a' && character <= 'z' {
				newText.WriteRune(shiftLetter(character, shift, false))
			} else {
				newText.WriteRune(character)
			}
		}
	} else {
		for _, character := range text {
			if character >= 'a' && character <= 'z' {
				newText.WriteRune(shiftLetter(character, shift, false))
			} else if character >= 'A' && character <= 'Z' {
				newText.WriteRune(shiftLetter(character, shift, true))
			} else {
				newText.WriteRune(character)
			}
		}
	}

	return newText.String()
}

// Decrypt takes an encrypted message and decrypts it using a shift.
// Uppercase can be preserved or skipped by using preserveUppercase.
// The function returns the shifted string, preserving each space or symbol
// as it was in the original message.
func Decrypt(text string, shift int, preserveUppercase bool) string {
	return Encrypt(text, -shift, preserveUppercase)
}

// shiftLetter shifts a letter by the given shift within the alphabet.
// The isUpper flag indicates whether the original letter is uppercase.
// It returns the shifted rune.
func shiftLetter(letter rune, shift int, isUpper bool) rune {
	initialPosition := fromAlphabetToPosition(letter, isUpper)
	finalPosition := ((initialPosition - 1 + shift) % 26 + 26) % 26 + 1
	shiftedLetter := fromPositionToAlphabet(finalPosition, isUpper)
	return shiftedLetter
}

// fromAlphabetToPosition returns the 1-based position of a letter in the alphabet.
// The isUpper flag indicates whether the letter is uppercase.
func fromAlphabetToPosition(letter rune, isUpper bool) int {
	if !isUpper {
		return int(letter - 'a') + 1
	}
	
	return int(letter - 'A') + 1
}

// fromPositionToAlphabet returns the letter corresponding to the given
// 1-based position in the alphabet. The toUpper flag controls the case of the output.
func fromPositionToAlphabet(position int, toUpper bool) rune {
	if !toUpper {
		return rune('a' + position - 1)
	}
	
	return rune('A' + position - 1)
}