package repository

import "errors"

var (
	ErrInvalidSubData = errors.New("Invalid subscription data")
	ErrNoSubIdExists = errors.New("No subscription with such id exists")
)
