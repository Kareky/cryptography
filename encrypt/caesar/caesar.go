package caesar

import (
	"strings"
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
)

// Encrypt takes a message and encrypts it, moving each character inside the sentence
// according to shift. a determines the alphabet used for the encryption.
// The function returns the shifted string, always lowercase,
// preserving each space or symbol as it was in the original message.
func Encrypt(text string, shift int, a alphabet.Alphabet) string {
	var encryptedStr strings.Builder
	for _, char := range text {
		char = unicode.ToLower(char)
		encryptedStr.WriteRune(alphabet.ShiftLetter(char, shift, a))
	}
	return encryptedStr.String()
}

// Decrypt takes an encrypted message and decrypts it using a shift.
// a determines the alphabet used for the encryption.
// The function returns the shifted string, always lowercase,
// preserving each space or symbol as it was in the original message.
func Decrypt(text string, shift int, a alphabet.Alphabet) string {
	return Encrypt(text, -shift, a)
}