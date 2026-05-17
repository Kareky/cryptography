// Package coordinates provides a Coordinate type representing a row‑column pair
// and functions to convert between different representations of coordinate lists.
package coordinates

import (
	"fmt"
	"strings"
)

// Coordinate is a row‑column pair.
type Coordinate struct {
	Row int
	Col int
}

// New returns a Coordinate with the given row and column.
func New(row int, col int) Coordinate {
	return Coordinate{Row: row, Col: col}
}

// ToString returns the concatenated string of the row and column digits.
// For example, Row=1,Col=2 yields "12".
func (c Coordinate) ToString() string {
	return fmt.Sprintf("%d%d", c.Row, c.Col)
}

// ToInt returns a two‑element slice [Row, Col].
func (c Coordinate) ToInt() []int {
	return []int{c.Row, c.Col}
}

// ToString converts a slice of Coordinates into a space‑separated string
// where each coordinate is represented as its row‑column digit string
// (e.g., [{1,2} {3,4}] → "12 34").
func ToString(c []Coordinate) string {
	var text strings.Builder
	for pos, coor := range c {
		text.WriteString(coor.ToString())
		if pos != len(c)-1 {
			text.WriteString(" ")
		}
	}
	return text.String()
}

// ToPairStrings converts a slice of Coordinates into a slice of strings,
// each being the concatenation of the coordinate's digits (e.g., "12").
func ToPairStrings(c []Coordinate) []string {
	var coordinates []string
	for _, coor := range c {
		coordinates = append(coordinates, coor.ToString())
	}
	return coordinates
}

// ToInt flattens a slice of Coordinates into a single slice of integers
// in the order Row, Col, Row, Col, …
func ToInt(c []Coordinate) []int {
	var coorL []int
	for _, coor := range c {
		coorL = append(coorL, coor.Row, coor.Col)
	}
	return coorL
}

// ToPairInts converts a slice of Coordinates into a slice of integer pairs,
// where each pair is a []int{Row, Col}.
func ToPairInts(c []Coordinate) [][]int {
	var pairs [][]int
	for _, coor := range c {
		pairs = append(pairs, coor.ToInt())
	}
	return pairs
}

// FromString parses a space‑separated string of concatenated row‑column
// digit pairs (e.g., "12 34") into a slice of Coordinates.
// It returns an error if a token does not consist of exactly two digits.
func FromString(strCoordinates string) ([]Coordinate, error) {
	coordinatePairs := strings.Fields(strCoordinates)
	return FromStringPairs(coordinatePairs)
}

// FromStringPairs converts a slice of two‑digit strings (e.g., ["12","34"])
// into a slice of Coordinates. Each string must be exactly two ASCII digits.
// It returns an error if any string violates these rules.
func FromStringPairs(coordinatePairs []string) ([]Coordinate, error) {
	c := []Coordinate{}
	for _, strCoordinate := range coordinatePairs {
		if len(strCoordinate) != 2 {
			return nil, ErrNotAPair
		}
		row, col := strCoordinate[0], strCoordinate[1]
		if !isDigit(row) || !isDigit(col) {
			return nil, ErrNotAnInt
		}
		c = append(c, Coordinate{Row: int(row - '0'), Col: int(col - '0')})
	}
	return c, nil
}

// FromInt converts a flat slice of integers (Row, Col, Row, Col, …)
// into a slice of Coordinates. It returns an error if the slice length is odd,
// meaning a coordinate is missing its row or column partner.
func FromInt(intCoordinates []int) ([]Coordinate, error) {
	c := []Coordinate{}
	if len(intCoordinates)%2 != 0 {
		return nil, ErrNotAPair
	}
	for i := 0; i < len(intCoordinates); i += 2 {
		c = append(c, Coordinate{Row: intCoordinates[i], Col: intCoordinates[i+1]})
	}
	return c, nil
}

// FromIntPairs converts a slice of integer pairs ([][]int) into a slice of Coordinates.
// Each inner slice must have exactly two elements.
func FromIntPairs(coordinatePairs [][]int) ([]Coordinate, error) {
	c := []Coordinate{}
	for _, coordinate := range coordinatePairs {
		if len(coordinate) != 2 {
			return nil, ErrNotAPair
		}
		c = append(c, Coordinate{Row: coordinate[0], Col: coordinate[1]})
	}
	return c, nil
}

// isDigit reports whether b is an ASCII digit ('0'–'9').
func isDigit(b byte) bool { return '0' <= b && b <= '9' }