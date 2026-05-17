// Package frequency provides functions to analyze the frequency of characters in a text.
package frequency

import (
	"sync"
	"unicode"

	"github.com/Kareky/cryptography/internal/alphabet"
)

// FreqResult holds letter frequencies for a specific alphabet.
type FreqResult struct {
    alphabet alphabet.Alphabet
    counts   []int
	once     sync.Once
    m        map[rune]int
}

// LetterFrequency counts how many times each letter of the given alphabet
// appears in text. It lowercases the text and ignores characters not in the alphabet.
// The returned slice has length a.Len(). It return FreqResult,
// which can be cast to the desired result.
// It panics if a is nil.
func LetterFrequency(text string, a alphabet.Alphabet) *FreqResult {
	if a == nil {
		panic("frequency: " + alphabet.ErrAlphabetNil.Error())
	}

	freq := make([]int, a.Len())
	for _, char := range text {
		char = unicode.ToLower(char)
		if pos, ok := a.Position(char); ok {
			freq[pos]++
		}
	}

	return &FreqResult{alphabet: a, counts: freq}
}

// Slice returns the frequency counts as a slice (index = position in alphabet).
func (fr *FreqResult) Slice() []int { return fr.counts }

// Map returns the frequency counts as a map[rune]int.
func (fr *FreqResult) Map() map[rune]int {
    fr.once.Do(func() {
        fr.m = make(map[rune]int, fr.alphabet.Len())
        for i := 0; i < fr.alphabet.Len(); i++ {
            fr.m[fr.alphabet.RuneFor(i)] = fr.counts[i]
        }
    })
    return fr.m
}