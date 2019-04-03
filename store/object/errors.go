package object

import "errors"

var (
	ErrInvalidKey = errors.New("invalid key, cannot be empty")
	ErrInvalidValue = errors.New("invalid value, cannot be empty or nil")
	ErrInvalidOutput = errors.New("invalid out, cannot be nil")
)

