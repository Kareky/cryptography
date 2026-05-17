package otp

import "errors"

var (
    ErrLengthMismatch 	= errors.New("pad length must equal the number of letters in the text")
	ErrEmptyPad			= errors.New("pad is empty or nil")
)