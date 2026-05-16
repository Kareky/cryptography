package random

import (
	"crypto/rand"
	"math/big"
	"unicode"

	"github.com/Kareky/cryptography/internal/alphabet"
)

// GeneratePad generates a sequence of cryptographically secure random integers.
// length specifies the number of integers to generate.
// max is the exclusive upper bound for each integer (i.e., values will be in [0, max)).
// It returns an error if max < 1, or if the random source fails.
func GeneratePad(length int, max int) ([]int, error) {
	return generatePad(length, max)
}

// GeneratePadWithAlphabet generates a sequence of cryptographically secure random integers.
// The number of integers is determined by length, and the exclusive upper bound for each
// integer is the length of the given alphabet (a.Len()). It returns an error if the
// random source fails.
func GeneratePadWithAlphabet(length int, a alphabet.Alphabet) ([]int, error) {
	return generatePad(length, a.Len())
}

// GeneratePadFromText generates a sequence of cryptographically secure random integers.
// The number of integers is determined by the number of characters in text that belong to a.
// Each integer is in the range [0, a.Len()). The text is lowercased before validation.
// It returns an error if the random source fails.
func GeneratePadFromText(text string, a alphabet.Alphabet) ([]int, error) {
	var length = 0
	for _, char := range text {
		char = unicode.ToLower(char)
		if _, ok := a.Position(char); ok {
			length++
		}
	}

	return generatePad(length, a.Len())
}

// generatePad contains the shared implementation for pad generation.
// It generates length random integers, each in the range [0, max). It returns an error
// if the random source fails.
func generatePad(length int, max int) ([]int, error) {
	if max < 1 {
		return nil, ErrMaxMustBePositive
	}

	seq := make([]int, 0, length)
	r := rand.Reader
	maxB := big.NewInt(int64(max))
	for range length {
		n, err := rand.Int(r, maxB)
		if err != nil {
			return nil, err
		}
		seq = append(seq, int(n.Int64()))
	}
	
	return seq, nil
}