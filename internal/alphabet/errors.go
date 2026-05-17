package alphabet

import "errors"

var (
    ErrAlphabetMismatch = errors.New("string is not compatible with the selected alphabet")
    ErrAlphabetNil      = errors.New("alphabet is nil")
)