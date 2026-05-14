package caesar

import (
	"strings"
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
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
				newText.WriteRune(alphabet.ShiftLetter(character, shift, false))
			} else {
				newText.WriteRune(character)
			}
		}
	} else {
		for _, character := range text {
			if character >= 'a' && character <= 'z' {
				newText.WriteRune(alphabet.ShiftLetter(character, shift, false))
			} else if character >= 'A' && character <= 'Z' {
				newText.WriteRune(alphabet.ShiftLetter(character, shift, true))
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