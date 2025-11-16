package repository

import (
	"context"
	"go-store-filter/internal/infrastructure/model"
)

type FilterRepository interface {
	Create(ctx context.Context, filter *model.Filter) error
	GetById(ctx context.Context, id int64) (*model.Filter, error)
	GetAll(ctx context.Context) ([]model.Filter, error)
	Update(ctx context.Context, filter *model.Filter) error
	Delete(ctx context.Context, id int64) error
}
