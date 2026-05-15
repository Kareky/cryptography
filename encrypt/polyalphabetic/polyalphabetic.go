package polyalphabetic

import (
	"strings"
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
)

// Encrypt takes a message and encrypts it, moving each character inside the sentence
// according to a cycle of shifts determined by the position in the alphabet of the letters
// of the key word passed.
// a determines the alphabet used for the cipher.
// The function returns the encrypted string, always lowercase, preserving each space or symbol
// as it was in the original message.
// It returns an error if encryptWord contains non-letters or letters that don't exist in the chosen alphabet.
// If encryptWord is empty, the text is returned unchanged.
func Encrypt(text, encryptWord string, a alphabet.Alphabet) (string, error) {
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

// toEncryptSequence transforms a key word into a slice of shift values,
// based on each letter's position in the alphabet.
// It returns an error if any character in the word is not in the alphabet.
func toEncryptSequence(encryptWord string, a alphabet.Alphabet) ([]int, error) {
	var encryptSequence []int
	for _, char := range encryptWord {
		char = unicode.ToLower(char)
		if pos, ok := a.Position(char); ok {
			encryptSequence = append(encryptSequence, pos)
		} else {
			return nil, alphabet.ErrAlphabetMismatch
		}
	}

	return encryptSequence, nil
}

// Decrypt takes a message and decrypts it, moving each character inside the sentence
// according to a cycle of shifts determined by the inverse of the position in the alphabet
// of the letters of the key word passed.
// a determines the alphabet used for the cipher.
// The function returns the decrypted string, always lowercase, preserving each space or symbol
// as it was in the original message.
// It returns an error if decryptWord contains non-letters or letters that don't exist in the chosen alphabet.
// If decryptWord is empty, the text is returned unchanged.
func Decrypt(text, decryptWord string, a alphabet.Alphabet) (string, error) {
	var encryptWord strings.Builder
	for _, char := range decryptWord {
		char = unicode.ToLower(char)
		pos, ok := a.Position(char)
		if !ok {
			return "", alphabet.ErrAlphabetMismatch
		}

		newPosition := (a.Len() - pos)%a.Len()
		char = a.RuneFor(newPosition)
		encryptWord.WriteRune(char)
	}

	return Encrypt(text, encryptWord.String(), a)
}