package otp

import (
	"strings"
	"unicode"

	"github.com/Kareky/cryptography/internal/alphabet"
	"github.com/Kareky/cryptography/internal/random"
)

// Encrypt encrypts text using the provided one‑time pad.
// The pad must contain exactly as many shift values as there are letters in text
// (spaces and punctuation are preserved and not counted). It returns an error if
// the pad length does not match the number of letters, or if the pad is empty.
// The ciphertext preserves the original formatting.
func Encrypt(text string, pad []int, a alphabet.Alphabet) (string, error) {
	err := checkLengthMismatch(text, pad, a)
	if err != nil {
		return "", err
	}

	return encrypt(text, pad, a, true), nil
}

// EncryptStripped encrypts text using the provided one‑time pad and removes
// all non‑letter characters from the output. It returns an error if
// the pad length does not match the number of letters, or if the pad is empty.
func EncryptStripped(text string, pad []int, a alphabet.Alphabet) (string, error) {
	err := checkLengthMismatch(text, pad, a)
	if err != nil {
		return "", err
	}

	return encrypt(text, pad, a, false), nil
}

// EncryptWithPad generates a cryptographically secure one‑time pad that
// matches the letters in text, then encrypts the text with it.
// The ciphertext preserves the original formatting. The generated pad is
// returned alongside the ciphertext so it can be used for decryption.
// It returns an error if the alphabet is nil, or if the random source fails.
func EncryptWithPad(text string, a alphabet.Alphabet) (string, []int, error) {
	pad, err := generatePadFromText(text, a)
	if err != nil {
		return "", nil, err
	}

	return encrypt(text, pad, a, true), pad, nil
}

// EncryptWithPadStripped generates a cryptographically secure one‑time pad that
// matches the letters in text, then encrypts the text with it and removes
// all non‑letter characters from the output. The generated pad is returned
// alongside the ciphertext.
// It returns an error if the alphabet is nil, or if the random source fails.
func EncryptWithPadStripped(text string, a alphabet.Alphabet) (string, []int, error) {
	pad, err := generatePadFromText(text, a)
	if err != nil {
		return "", nil, err
	}

	return encrypt(text, pad, a, false), pad, nil
}

// Decrypt decrypts a ciphertext that was encrypted with the given one‑time pad.
// The pad must be identical to the one used during encryption (length and values),
// and must contain exactly as many shift values as there are letters in the
// ciphertext. It returns an error if the pad is empty, the alphabet is nil,
// or the pad length does not match the number of letters. The decrypted text
// preserves all spaces and punctuation from the ciphertext.
func Decrypt(text string, pad []int, a alphabet.Alphabet) (string, error) {
	err := checkLengthMismatch(text, pad, a)
	if err != nil {
		return "", err
	}

	return decrypt(text, pad, a), nil
}

// generatePadFromText returns a cryptographically secure pad whose length
// equals the number of alphabet letters in text. It returns an error if
// the alphabet is nil or if the random source fails.
func generatePadFromText(text string, a alphabet.Alphabet) ([]int, error) {
	if a == nil {
		return nil, alphabet.ErrAlphabetNil
	}

	pad, err := random.GeneratePadFromText(text, a)
	if err != nil {
		return nil, err
	}

	return pad, nil
}

// checkLengthMismatch verifies that the number of shift‑able letters in text
// equals len(pad). It returns ErrEmptyPad if the pad is empty,
// ErrAlphabetNil if the alphabet is nil, or ErrLengthMismatch if the
// counts do not match.
func checkLengthMismatch(text string, pad []int, a alphabet.Alphabet) error {
	if len(pad) == 0 {
		return ErrEmptyPad
	}

	if a == nil {
		return alphabet.ErrAlphabetNil
	}

	var i = 0
	for _, char := range text {
		char = unicode.ToLower(char)
		if _, ok := a.Position(char); ok {
			i++
		}
	}

	if i != len(pad) {
		return ErrLengthMismatch
	}

	return nil
}

// encrypt contains the core encryption logic shared by all Encrypt variants.
// It shifts each letter by the corresponding pad value; non‑letter characters
// are either preserved (if preserveFormatting is true) or dropped.
// The caller must guarantee that the pad length matches the number of letters.
func encrypt(text string, pad []int, a alphabet.Alphabet, preserveFormatting bool) string {
	var encryptedStr strings.Builder
	var i = 0
	for _, char := range text {
		char = unicode.ToLower(char)
		if _, ok := a.Position(char); ok {
			char = alphabet.ShiftLetter(char, pad[i], a)
			encryptedStr.WriteRune(char)
			i++
		} else if preserveFormatting {
			encryptedStr.WriteRune(char)
		}
	}

	return encryptedStr.String()
}

// decrypt contains the core decryption logic. It shifts each letter by the
// inverse of the corresponding pad value, effectively reversing the encryption.
// The caller must guarantee that the pad length matches the number of letters.
// Non‑letter characters are always preserved, assuming that, if they needed to be stripped
// they were stripped at the encryption call level.
func decrypt(text string, pad []int, a alphabet.Alphabet) string {
	var encryptedStr strings.Builder
	var i = 0
	for _, char := range text {
		char = unicode.ToLower(char)
		if _, ok := a.Position(char); ok {
			char = alphabet.ShiftLetter(char, -pad[i], a)
			i++
		}
		encryptedStr.WriteRune(char)
	}

	return encryptedStr.String()
}