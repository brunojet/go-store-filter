package usecase

import (
	"context"
	"go-store-filter/internal/infrastructure/model"
	"go-store-filter/internal/repository"
)

type FilterTypeService struct {
	repo repository.FilterTypeRepository
}

func NewFilterTypeService(repo repository.FilterTypeRepository) *FilterTypeService {
	return &FilterTypeService{repo: repo}
}

func (s *FilterTypeService) Create(ctx context.Context, filterType *model.FilterType) error {
	return s.repo.Create(ctx, filterType)
}

func (s *FilterTypeService) GetById(ctx context.Context, id int64) (*model.FilterType, error) {
	return s.repo.GetById(ctx, id)
}

func (s *FilterTypeService) GetAll(ctx context.Context) ([]model.FilterType, error) {
	return s.repo.GetAll(ctx)
}

func (s *FilterTypeService) Update(ctx context.Context, filterType *model.FilterType) error {
	return s.repo.Update(ctx, filterType)
}

func (s *FilterTypeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
