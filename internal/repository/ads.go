package repository

import (
	"context"
	"subs-service/internal/domain"

	"github.com/google/uuid"
)

type SubsRepo interface {
	GetSub(ctx context.Context, id uuid.UUID) (*domain.Sub, error)
	PostSub(ctx context.Context, sub *domain.Sub) (*domain.Sub, error)
	PutSub(ctx context.Context, id uuid.UUID, sub *domain.Sub) (*domain.Sub, error)
	DeleteSub(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	ListSubs(ctx context.Context, opts domain.FilterOpts) ([]*domain.Sub, error)
	GetSummary(ctx context.Context, opts domain.FilterOpts) (*domain.Summary, error)
}
