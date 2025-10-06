package repository

import "errors"

var (
	ErrInvalidSubData = errors.New("invalid subscription data")
	ErrNoSubIDExists = errors.New("no subscription with such id exists")
)
