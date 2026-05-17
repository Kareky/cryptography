package coordinates

import (
	"errors"
)

var (
	ErrNotAnInt		= errors.New("trying to convert a coordinate that wasn't an int")
	ErrNotAPair		= errors.New("a set of coordinate was not sent in pairs")
)