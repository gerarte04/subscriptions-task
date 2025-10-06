package database

import "errors"

var (
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrUniqueViolation = errors.New("unique violation")
	ErrCheckViolation = errors.New("check violation")
	ErrUndocumented = errors.New("undocumented database error")
)
