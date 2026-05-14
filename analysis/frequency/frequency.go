// Package frequency provides functions to analyze the frequency of characters in a text.
package frequency

import (
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
)


func LetterFrequencyLatin(text string, alphabet alphabet.Alphabet) []int {
	freq := make([]int, alphabet.Len())
	for _, char := range text {
		char = unicode.ToLower(char)
		if pos, ok := alphabet.Position(char); ok {
			freq[pos]++
		}
	}

	return freq
}