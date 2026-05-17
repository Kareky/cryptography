package polybius

import (
	"strings"
	"unicode"

	"github.com/Kareky/cryptography/internal/coordinates"
)

// Encrypt encrypts text using the Polybius square pSquare.
// It lowercases each character and converts it to its coordinate in the square.
// Characters not present in the square (including spaces and punctuation) are
// silently skipped. It returns an error if pSquare is nil.
func Encrypt(text string, pSquare Square) ([]coordinates.Coordinate, error) {
	if pSquare == nil {
		return nil, ErrSquareNil
	}

	var coordinates = []coordinates.Coordinate{}
	for _, char := range text {
		char = unicode.ToLower(char)
		if coor, ok := pSquare.Coordinate(char); ok {
			coordinates = append(coordinates, coor)
		}
	}

	return coordinates, nil
}

// Decrypt converts a slice of coordinates back into a string using the
// Polybius square pSquare. Coordinates that do not exist in the square are
// silently skipped. It returns an error if pSquare is nil.
func Decrypt(coordinates []coordinates.Coordinate, pSquare Square) (string, error) {
	if pSquare == nil {
		return "", ErrSquareNil
	}

	var text strings.Builder
	for _, coor := range coordinates {
		if char, ok := pSquare.RuneAt(coor); ok {
			text.WriteRune(char)
		}
	}

	return text.String(), nil
}