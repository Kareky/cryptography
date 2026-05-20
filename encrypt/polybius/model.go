// Package polybius implements the Polybius square cipher.
// A Square maps letters to row‑column coordinates and vice versa,
// and supports aliases to map missing letters (e.g., 'j'→'i') into the square.
package polybius

import (
	"maps"
	"math"
	"unicode"

	"github.com/Kareky/cryptography/internal/alphabet"
	"github.com/Kareky/cryptography/internal/coordinates"
)

// Square is the interface for a Polybius square.
// It can translate a rune to its coordinate, or a coordinate back to its rune.
type Square interface {
	CellCount() int
	SideLength() int
	Coordinate(char rune) (coordinates.Coordinate, bool)
	RuneAt(c coordinates.Coordinate) (rune, bool)
	Name() string
}

// SquareImp is the default implementation of a Polybius square.
// It holds a mapping of coordinates to runes, plus an aliases map that
// can redirect an input rune (e.g., 'j'→'i') before lookup.
type SquareImp struct {
	mapping map[coordinates.Coordinate]rune
	aliases	map[rune]rune
	name    string
}

func (s SquareImp) Name() string    { return s.name }
func (s SquareImp) CellCount() int  { return len(s.mapping) }
func (s SquareImp) SideLength() int { return int(math.Sqrt(float64(s.CellCount()))) }

// RuneAt returns the rune at the given coordinate.
// It returns false if the coordinate does not exist in the square.
func (s SquareImp) RuneAt(c coordinates.Coordinate) (rune, bool) {
	char, ok := s.mapping[c]
	if !ok {
		return -1, false
	}

	return char, true
}

// Coordinate returns the coordinate of char in the square.
// If char is not found directly, aliases are consulted first.
// It returns false if the character is absent even after alias lookup.
func (s SquareImp) Coordinate(char rune) (coordinates.Coordinate, bool) {
	if alias, ok := s.aliases[char]; ok{
		char = alias
	}

	for coordinate, c := range s.mapping {
		if char == c {
			return coordinate, true
		}
	}

	return coordinates.Coordinate{}, false
}

// NewSquare creates a validated SquareImp.
// It returns an error if the mapping does not form a perfect square,
// if any coordinate from (0,0) to (side‑1,side‑1) is missing,
// or if any rune is not lowercase.
// The aliases map is stored as‑is; it is not checked for validity.
func NewSquare(name string, mapping map[coordinates.Coordinate]rune, aliases map[rune]rune) (*SquareImp, error) {
	totalCells := len(mapping)
    side := int(math.Sqrt(float64(totalCells)))
    if side*side != totalCells {
        return nil, ErrNoPerfectSquare
    }

	for r := range side {
        for c := range side {
            if char, exists := mapping[coordinates.New(r, c)]; !exists {
                return nil, ErrMissingCoordinate
            } else if !unicode.IsLower(char) && unicode.IsLetter(char) {
				return nil, ErrMissingLowerCase(coordinates.New(r, c))
			}
        }
    }

	return &SquareImp{
		name: name,
		mapping: mapping,
		aliases: aliases,
	}, nil
}

// GridToMap returns a coordinate‑to‑rune mapping from a 2‑D rune grid.
// It rejects cells with no runes or uppercase letters, but accepts symbols.
func GridToMap(grid [][]rune) (map[coordinates.Coordinate]rune, error) {
	var mapping = map[coordinates.Coordinate]rune{}
	for r, rows := range grid {
		for c, char := range rows {
			if char == 0 {
				return nil, ErrMissingCoordinate
			} else if !unicode.IsLower(char) && unicode.IsLetter(char) {
				return nil, ErrMissingLowerCase(coordinates.New(r, c))
			}
			mapping[coordinates.New(r, c)] = char
		}
	}

	return mapping, nil
}

// LatinNoZSquare is a 5×5 square containing a‑y; 'z' is aliased to 's'.
var LatinNoZSquare, _ = NewSquare("LatinNoZ", LatinNoZ, LatinNoZAliases)

// LatinNoJSquare is a 5×5 square containing a‑z except 'j' (merged with 'i').
var LatinNoJSquare, _ = NewSquare("LatinNoJ", LatinNoJ, LatinNoJAliases)

// LatinNoZ is the mapping for a 5×5 square with letters a‑y.
var LatinNoZ, _ = GridToMap([][]rune{
	{'a','b','c','d','e'},
	{'f','g','h','i','j'},
	{'k','l','m','n','o'},
	{'p','q','r','s','t'},
	{'u','v','w','x','y'},
})
/*var LatinNoZ = map[coordinates.Coordinate]rune {
	coordinates.New(0, 0): 'a',
	coordinates.New(0, 1): 'b',
	coordinates.New(0, 2): 'c',
	coordinates.New(0, 3): 'd',
	coordinates.New(0, 4): 'e',
		coordinates.New(1, 0): 'f',
		coordinates.New(1, 1): 'g',
		coordinates.New(1, 2): 'h',
		coordinates.New(1, 3): 'i',
		coordinates.New(1, 4): 'j',
			coordinates.New(2, 0): 'k',
			coordinates.New(2, 1): 'l',
			coordinates.New(2, 2): 'm',
			coordinates.New(2, 3): 'n',
			coordinates.New(2, 4): 'o',
				coordinates.New(3, 0): 'p',
				coordinates.New(3, 1): 'q',
				coordinates.New(3, 2): 'r',
				coordinates.New(3, 3): 's',
				coordinates.New(3, 4): 't',
					coordinates.New(4, 0): 'u',
					coordinates.New(4, 1): 'v',
					coordinates.New(4, 2): 'w',
					coordinates.New(4, 3): 'x',
					coordinates.New(4, 4): 'y',
}*/

// LatinNoJ is the mapping for a 5×5 square with letters a‑z, merging 'j' into 'i'.
var LatinNoJ, _ = GridToMap([][]rune{
	{'a','b','c','d','e'},
	{'f','g','h','i','k'},
	{'l','m','n','o','p'},
	{'q','r','s','t','u'},
	{'v','w','x','y','z'},
})
/*var LatinNoJ = map[coordinates.Coordinate]rune {
	coordinates.New(0, 0): 'a',
	coordinates.New(0, 1): 'b',
	coordinates.New(0, 2): 'c',
	coordinates.New(0, 3): 'd',
	coordinates.New(0, 4): 'e',
		coordinates.New(1, 0): 'f',
		coordinates.New(1, 1): 'g',
		coordinates.New(1, 2): 'h',
		coordinates.New(1, 3): 'i',
		coordinates.New(1, 4): 'k',
			coordinates.New(2, 0): 'l',
			coordinates.New(2, 1): 'm',
			coordinates.New(2, 2): 'n',
			coordinates.New(2, 3): 'o',
			coordinates.New(2, 4): 'p',
				coordinates.New(3, 0): 'q',
				coordinates.New(3, 1): 'r',
				coordinates.New(3, 2): 's',
				coordinates.New(3, 3): 't',
				coordinates.New(3, 4): 'u',
					coordinates.New(4, 0): 'v',
					coordinates.New(4, 1): 'w',
					coordinates.New(4, 2): 'x',
					coordinates.New(4, 3): 'y',
					coordinates.New(4, 4): 'z',
}*/

// LatinNoZAliases maps 'z' to 's' and includes all common accents.
var LatinNoZAliases = func() map[rune]rune {
	var aliases = map[rune]rune {'z':'s'}
    return NewAliasMapWithCommonAccents(aliases)
}()

// LatinNoJAliases maps 'j' to 'i' and includes all common accents.
var LatinNoJAliases = func() map[rune]rune {
    var aliases = map[rune]rune {'j':'i'}
    return NewAliasMapWithCommonAccents(aliases)
}()

// NewAliasMapWithCommonAccents creates a map that combines the common
// accent‑stripping map with additional custom aliases.
// Later aliases overwrite earlier ones in case of key conflicts.
func NewAliasMapWithCommonAccents(additionalAliases ...map[rune]rune) map[rune]rune {
	m := make(map[rune]rune, len(alphabet.CommonAccents)+len(additionalAliases))
    maps.Copy(m, alphabet.CommonAccents)

	for _, aliases := range additionalAliases {
		maps.Copy(m, aliases)
	}

	return m
}