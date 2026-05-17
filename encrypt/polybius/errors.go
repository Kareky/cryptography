package polybius

import (
	"errors"
	"fmt"

	"github.com/Kareky/cryptography/internal/coordinates"
)

var (
	ErrSquareNil			= errors.New("polybius.square is nil")
	ErrNoPerfectSquare		= errors.New("mapping does not form a perfect square")
	ErrMissingCoordinate	= errors.New("mapping is missing coordinate")
)

func ErrMissingLowerCase(c coordinates.Coordinate) error {
	return fmt.Errorf("rune at coordinate %v is not lowercase", c)
}