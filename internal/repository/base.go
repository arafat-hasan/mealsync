package repository

import (
	"context"

	"gorm.io/gorm"
)

// baseRepository implements common CRUD operations
type baseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new instance of baseRepository
func NewBaseRepository[T any](db *gorm.DB) *baseRepository[T] {
	return &baseRepository[T]{db: db}
}

// Create creates a new entity
func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// FindByID finds an entity by ID
func (r *baseRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll finds all entities
func (r *baseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// FindActive finds active entities based on conditions
func (r *baseRepository[T]) FindActive(ctx context.Context, conditions map[string]interface{}) ([]T, error) {
	var entities []T
	query := r.db.WithContext(ctx).Where("is_active = ?", true)

	for key, value := range conditions {
		query = query.Where(key+" = ?", value)
	}

	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// Update updates an entity
func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

// Delete soft deletes an entity
func (r *baseRepository[T]) Delete(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Delete(entity).Error
}

// HardDelete permanently deletes an entity
func (r *baseRepository[T]) HardDelete(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Unscoped().Delete(entity).Error
}

// WithTransaction executes a function within a transaction
func (r *baseRepository[T]) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
