package postgres

import (
	"context"
	"errors"
	"fmt"
	"subs-service/internal/domain"
	"subs-service/internal/repository"
	"subs-service/pkg/database"
	pkgPostgres "subs-service/pkg/database/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubsRepo struct {
	pool *pgxpool.Pool
}

func NewSubsRepo(pool *pgxpool.Pool) *SubsRepo {
	return &SubsRepo{
		pool: pool,
	}
}

func (r *SubsRepo) GetSub(ctx context.Context, id uuid.UUID) (*domain.Sub, error) {
	const op = "SubsRepo.GetSub"

	query := `SELECT * FROM subs WHERE id = $1`

	var sub domain.Sub
	err := r.pool.QueryRow(
		ctx, query, id,
	).Scan(&sub.Id, &sub.UserId, &sub.ServiceName, &sub.Price, &sub.StartDate, &sub.EndDate)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNoSubIdExists)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sub, nil
}

func (r *SubsRepo) PostSub(ctx context.Context, sub *domain.Sub) (uuid.UUID, error) {
	const op = "SubsRepo.PostSub"

	query := 
		`INSERT INTO subs (user_id, service_name, price, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var subId uuid.UUID
	err := r.pool.QueryRow(
		ctx, query,
		sub.UserId, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate,
	).Scan(&subId)

	if err != nil {
		pgErr := pkgPostgres.DetectError(err)

		if errors.Is(pgErr, database.ErrCheckViolation) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrInvalidSubData)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return subId, nil
}

func (r *SubsRepo) PutSub(ctx context.Context, id uuid.UUID, sub *domain.Sub) error {
	const op = "SubsRepo.PutSub"

	query :=
		`UPDATE subs SET user_id = $1, service_name = $2, price = $3, start_date = $4, end_date = $5
			WHERE id = $6`

	tag, err := r.pool.Exec(
		ctx, query,
		sub.UserId, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate, id,
	)

	if err != nil {
		pgErr := pkgPostgres.DetectError(err)

		if errors.Is(pgErr, database.ErrCheckViolation) {
			return fmt.Errorf("%s: %w", op, repository.ErrInvalidSubData)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	if tag.RowsAffected() != 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrNoSubIdExists)
	}

	return nil
}

func (r *SubsRepo) DeleteSub(ctx context.Context, id uuid.UUID) error {
	const op = "SubsRepo.DeleteSub"

	query := "DELETE FROM subs WHERE id = $1"

	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if tag.RowsAffected() != 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrNoSubIdExists)
	}

	return nil
}

func (r *SubsRepo) ListSubs(ctx context.Context, opts domain.FilterOpts) ([]*domain.Sub, error) {
	const op = "SubRepo.ListSubs"

	query := "SELECT * FROM subs WHERE user_id = $1"

	if opts.PageToken != uuid.Nil {
		query = fmt.Sprintf("%s AND id > $2", query)
	}

	if len(opts.ServiceName) != 0 {
		query = fmt.Sprintf("%s AND service_name = $3", query)
	}

	if opts.PageSize != 0 {
		query = fmt.Sprintf("%s LIMIT $4", query)
	}

	rows, err := r.pool.Query(
		ctx, query,
		opts.UserId, opts.PageToken, opts.ServiceName, opts.PageSize,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	subs := []*domain.Sub{}

	for rows.Next() {
		var sub domain.Sub

		if err := rows.Scan(
			&sub.Id, &sub.UserId, &sub.ServiceName, &sub.Price, &sub.StartDate, &sub.EndDate,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		subs = append(subs, &sub)
	}

	return subs, nil
}

func (r *SubsRepo) GetSummary(ctx context.Context, opts domain.FilterOpts) (*domain.Summary, error) {
	const op = "SubRepo.GetSummary"

	query := "SELECT SUM(price) FROM subs WHERE user_id = $1"

	if len(opts.ServiceName) != 0 {
		query = fmt.Sprintf("%s AND service_name = $2", query)
	}

	var sum domain.Summary
	err := r.pool.QueryRow(
		ctx, query,
		opts.UserId, opts.ServiceName,
	).Scan(&sum.TotalPrice)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sum, nil
}
