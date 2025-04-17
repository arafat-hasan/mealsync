package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// menuSetRepository implements MenuSetRepository interface
type menuSetRepository struct {
	*baseRepository[model.MenuSet]
	db *gorm.DB
}

// NewMenuSetRepository creates a new instance of MenuSetRepository
func NewMenuSetRepository(db *gorm.DB) MenuSetRepository {
	return &menuSetRepository{
		baseRepository: NewBaseRepository[model.MenuSet](db),
		db:             db,
	}
}

// Create creates a new menu set
func (r *menuSetRepository) Create(ctx context.Context, menuSet *model.MenuSet) error {
	return r.baseRepository.Create(ctx, menuSet)
}

// FindByID finds a menu set by ID
func (r *menuSetRepository) FindByID(ctx context.Context, id uint) (*model.MenuSet, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all menu sets
func (r *menuSetRepository) FindAll(ctx context.Context) ([]model.MenuSet, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active menu sets based on conditions
func (r *menuSetRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuSet, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a menu set
func (r *menuSetRepository) Update(ctx context.Context, menuSet *model.MenuSet) error {
	return r.baseRepository.Update(ctx, menuSet)
}

// Delete soft deletes a menu set
func (r *menuSetRepository) Delete(ctx context.Context, menuSet *model.MenuSet) error {
	return r.baseRepository.Delete(ctx, menuSet)
}

// HardDelete permanently deletes a menu set
func (r *menuSetRepository) HardDelete(ctx context.Context, menuSet *model.MenuSet) error {
	return r.baseRepository.HardDelete(ctx, menuSet)
}

// AddMenuItem adds a menu item to the menu set
func (r *menuSetRepository) AddMenuItem(ctx context.Context, setItem *model.MenuSetItem) error {
	return r.db.WithContext(ctx).Create(setItem).Error
}

// RemoveMenuItem removes a menu item from the menu set
func (r *menuSetRepository) RemoveMenuItem(ctx context.Context, setItem *model.MenuSetItem) error {
	return r.db.WithContext(ctx).Delete(setItem).Error
}

// FindMenuItems finds all menu items in a menu set
func (r *menuSetRepository) FindMenuItems(ctx context.Context, menuSetID uint) ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := r.db.WithContext(ctx).
		Joins("JOIN menu_set_items ON menu_set_items.menu_item_id = menu_items.id").
		Where("menu_set_items.menu_set_id = ?", menuSetID).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
