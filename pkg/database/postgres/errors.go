package postgres

import (
	"errors"
	"subs-service/pkg/database"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	PgForeignKeyViolation = "23503"
	PgUniqueViolation = "23505"
	PgCheckViolation = "23514"
)

var (
	codeErrors = map[string]error {
		PgForeignKeyViolation: database.ErrForeignKeyViolation,
		PgUniqueViolation: database.ErrUniqueViolation,
		PgCheckViolation: database.ErrCheckViolation,
	}
)

func DetectError(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return database.ErrUndocumented
	}

	if dbErr, ok := codeErrors[pgErr.Code]; ok {
		return dbErr
	}

	return database.ErrUndocumented
}
