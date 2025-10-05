package service

import (
	"context"
	"fmt"
	"subs-service/internal/domain"
	"subs-service/internal/repository"

	"github.com/google/uuid"
)

type SubService struct {
	subRepo repository.SubsRepo
}

func NewSubService(subRepo repository.SubsRepo) *SubService {
	return &SubService{
		subRepo: subRepo,
	}
}

func (s *SubService) GetSub(ctx context.Context, id uuid.UUID) (*domain.Sub, error) {
	const op = "SubService.GetSub"

	sub, err := s.subRepo.GetSub(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return sub, nil
}

func (s *SubService) PostSub(ctx context.Context, sub *domain.Sub) (*domain.Sub, error) {
	const op = "SubService.PostSub"

	id, err := s.subRepo.PostSub(ctx, sub)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sub.Id = id

	return sub, nil
}

func (s *SubService) PutSub(ctx context.Context, id uuid.UUID, sub *domain.Sub) (*domain.Sub, error) {
	const op = "SubService.PutSub"

	if err := s.subRepo.PutSub(ctx, id, sub); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sub.Id = id

	return sub, nil
}

func (s *SubService) DeleteSub(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	const op = "SubService.DeleteSub"

	if err := s.subRepo.DeleteSub(ctx, id); err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *SubService) ListSubs(ctx context.Context, opts domain.FilterOpts) ([]*domain.Sub, error) {
	const op = "SubService.ListSubs"

	subs, err := s.subRepo.ListSubs(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subs, nil
}

func (s *SubService) GetSummary(ctx context.Context, opts domain.FilterOpts) (*domain.Summary, error) {
	const op = "SubService.GetSummary"

	sum, err := s.subRepo.GetSummary(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return sum, nil
}
