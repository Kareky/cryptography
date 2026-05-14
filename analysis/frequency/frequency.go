// Package frequency provides functions to analyze the frequency of characters in a text.
package frequency

import (
	"unicode"
	"github.com/Kareky/cryptography/internal/alphabet"
)


func LetterFrequency(text string, alphabet alphabet.Alphabet) []int {
	freq := make([]int, alphabet.Len())
	for _, char := range text {
		char = unicode.ToLower(char)
		if pos, ok := alphabet.Position(char); ok {
			freq[pos]++
		}
	}

	return freq
}

func FromArrayToMap(freq []int, a alphabet.Alphabet) map[rune]int {
	if len(freq) != a.Len() {
		return nil
	}

	freqMap := make(map[rune]int, a.Len())
	for i := 0; i < a.Len(); i++ {
		freqMap[a.RuneFor(i)] = freq[i]
	}

	return freqMap
}