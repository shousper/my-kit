package raw

import "errors"

var (
	ErrNotFound     = errors.New("entry not found")
	ErrInvalidKey   = errors.New("invalid key, cannot be empty")
	ErrInvalidValue = errors.New("invalid value, cannot be empty or nil")
)
