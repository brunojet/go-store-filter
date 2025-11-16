package repository

import (
	"context"
	"go-store-filter/internal/infrastructure/model"
)

type FilterTypeRepository interface {
	Create(ctx context.Context, filterType *model.FilterType) error
	GetById(ctx context.Context, id int64) (*model.FilterType, error)
	GetAll(ctx context.Context) ([]model.FilterType, error)
	Update(ctx context.Context, filterType *model.FilterType) error
	Delete(ctx context.Context, id int64) error
}
