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
	return encrypt(text, shift, a, true)
}

// EncryptStripped takes a message and encrypts it, moving each character inside the sentence
// according to shift. a determines the alphabet used for the encryption.
// The function returns the shifted string, always lowercase,
// with all spaces and symbols removed.
func EncryptStripped(text string, shift int, a alphabet.Alphabet) string {
	return encrypt(text, shift, a, false)
}

// encrypt contains the core logic of encryption as described in Encrypt and EncryptStripped.
// preserveFormatting is a boolean that decide if formatting will be kept during characters shifting,
// or if they will be removed.
func encrypt(text string, shift int, a alphabet.Alphabet, preserveFormatting bool) string {
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