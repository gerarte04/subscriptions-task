package database

import "errors"

var (
	ErrForeignKeyViolation = errors.New("Foreign key violation")
	ErrUniqueViolation = errors.New("Unique violation")
	ErrCheckViolation = errors.New("Check violation")
	ErrUndocumented = errors.New("Undocumented database error")
)
