package repository

import (
	"context"

	"gorm.io/gorm"
)

// baseRepository implements common CRUD operations
type baseRepository[T any] struct {
	DB *gorm.DB
}

// NewBaseRepository creates a new base repository
func NewBaseRepository[T any](db *gorm.DB) *baseRepository[T] {
	return &baseRepository[T]{DB: db}
}

// Create implements the Create operation
func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.DB.WithContext(ctx).Create(entity).Error
}

// FindByID implements the FindByID operation
func (r *baseRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.DB.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements the FindAll operation
func (r *baseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.DB.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// Update implements the Update operation
func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.DB.WithContext(ctx).Save(entity).Error
}

// Delete implements the Delete operation
func (r *baseRepository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return r.DB.WithContext(ctx).Delete(&entity, id).Error
}
