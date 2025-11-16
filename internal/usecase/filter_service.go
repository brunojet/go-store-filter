package usecase

import (
	"context"
	"go-store-filter/internal/infrastructure/model"
	"go-store-filter/internal/repository"
)

type FilterService struct {
	repo repository.FilterRepository
}

func NewFilterService(repo repository.FilterRepository) *FilterService {
	return &FilterService{repo: repo}
}

func (s *FilterService) Create(ctx context.Context, filter *model.Filter) error {
	return s.repo.Create(ctx, filter)
}

func (s *FilterService) GetById(ctx context.Context, id int64) (*model.Filter, error) {
	return s.repo.GetById(ctx, id)
}

func (s *FilterService) GetAll(ctx context.Context) ([]model.Filter, error) {
	return s.repo.GetAll(ctx)
}

func (s *FilterService) Update(ctx context.Context, filter *model.Filter) error {
	return s.repo.Update(ctx, filter)
}

func (s *FilterService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
