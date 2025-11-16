package repository

import (
	"context"
	"go-store-filter/internal/infrastructure/model"

	"gorm.io/gorm"
)

type GormFilterRepository struct {
	db *gorm.DB
}

func NewGormFilterRepository(db *gorm.DB) *GormFilterRepository {
	return &GormFilterRepository{db: db}
}

func (r *GormFilterRepository) Create(ctx context.Context, filter *model.Filter) error {
	return r.db.WithContext(ctx).Create(filter).Error
}

func (r *GormFilterRepository) GetById(ctx context.Context, id int64) (*model.Filter, error) {
	var filter model.Filter
	err := r.db.WithContext(ctx).First(&filter, id).Error
	if err != nil {
		return nil, err
	}
	return &filter, nil
}

func (r *GormFilterRepository) GetAll(ctx context.Context) ([]model.Filter, error) {
	var filters []model.Filter
	err := r.db.WithContext(ctx).Find(&filters).Error
	return filters, err
}

func (r *GormFilterRepository) Update(ctx context.Context, filter *model.Filter) error {
	return r.db.WithContext(ctx).Save(filter).Error
}

func (r *GormFilterRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Filter{}, id).Error
}
