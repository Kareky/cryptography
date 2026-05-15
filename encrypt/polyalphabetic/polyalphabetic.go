package polyalphabetic

import (
	"strings"
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
)

// Encrypt takes a message and encrypts it, moving each character inside the sentence
// according to a cicle of shifts determined by position in the alphabet of the letters of the word passed.
// a determines the alphabet used for the encryption.
// The function returns the shifted string, always lowercase, preserving each space or symbol as it was in the original message.
// It return error if encryptWord contain non-letters or letters that doesn't exists in the chosen alphabet.
// If the encryptWord is empty, the text is returned unchanged
func Encrypt(text string, encryptWord string, a alphabet.Alphabet) (string, error) {
	if len(encryptWord) == 0 {
		return text, nil
	}

	var encryptedStr strings.Builder
	var i = 0
	
	encryptSeq, err := toEncryptSequence(encryptWord, a)
	if err != nil {
		return "", err
	}

	for _, char := range text {
		char = unicode.ToLower(char)
		if _, ok := a.Position(char); ok {
			char = alphabet.ShiftLetter(char, encryptSeq[i % len(encryptSeq)], a)
			i++
		}
		encryptedStr.WriteRune(char)
	}

	return encryptedStr.String(), nil
}

// toEncryptSequence transform an encrypt word into a slice of numbers
// it return error if the characters into the word aren't known inside the alphabet
func toEncryptSequence(encryptWord string, a alphabet.Alphabet) ([]int, error) {
	var encryptSequence []int
	for _, char := range encryptWord {
		if pos, ok := a.Position(char); ok {
			encryptSequence = append(encryptSequence, pos)
		} else {
			return nil, alphabet.ErrAlphabetMismatch
		}
	}

	return encryptSequence, nil
}