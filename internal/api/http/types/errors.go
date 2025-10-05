package types

import "errors"

var (
	ErrBadPriceValue = errors.New("bad price value, must be positive and less than max")
	ErrBadServiceNameLength = errors.New("bad service name length (must be non zero and less than max)")
)
