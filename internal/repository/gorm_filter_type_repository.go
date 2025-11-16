package repository

import (
	"context"
	"go-store-filter/internal/infrastructure/model"

	"gorm.io/gorm"
)

type GormFilterTypeRepository struct {
	db *gorm.DB
}

func NewGormFilterTypeRepository(db *gorm.DB) *GormFilterTypeRepository {
	return &GormFilterTypeRepository{db: db}
}

func (r *GormFilterTypeRepository) Create(ctx context.Context, filterType *model.FilterType) error {
	return r.db.WithContext(ctx).Create(filterType).Error
}

func (r *GormFilterTypeRepository) GetById(ctx context.Context, id int64) (*model.FilterType, error) {
	var filterType model.FilterType
	err := r.db.WithContext(ctx).First(&filterType, id).Error
	if err != nil {
		return nil, err
	}
	return &filterType, nil
}

func (r *GormFilterTypeRepository) GetAll(ctx context.Context) ([]model.FilterType, error) {
	var filterTypes []model.FilterType
	err := r.db.WithContext(ctx).Find(&filterTypes).Error
	return filterTypes, err
}

func (r *GormFilterTypeRepository) Update(ctx context.Context, filterType *model.FilterType) error {
	return r.db.WithContext(ctx).Save(filterType).Error
}

func (r *GormFilterTypeRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.FilterType{}, id).Error
}
