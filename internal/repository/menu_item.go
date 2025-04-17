package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// menuItemRepository implements MenuItemRepository interface
type menuItemRepository struct {
	*baseRepository[model.MenuItem]
	db *gorm.DB
}

// NewMenuItemRepository creates a new instance of MenuItemRepository
func NewMenuItemRepository(db *gorm.DB) MenuItemRepository {
	return &menuItemRepository{
		baseRepository: NewBaseRepository[model.MenuItem](db),
		db:             db,
	}
}

// Create creates a new menu item
func (r *menuItemRepository) Create(ctx context.Context, item *model.MenuItem) error {
	return r.baseRepository.Create(ctx, item)
}

// FindByID finds a menu item by ID
func (r *menuItemRepository) FindByID(ctx context.Context, id uint) (*model.MenuItem, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all menu items
func (r *menuItemRepository) FindAll(ctx context.Context) ([]model.MenuItem, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active menu items based on conditions
func (r *menuItemRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuItem, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a menu item
func (r *menuItemRepository) Update(ctx context.Context, item *model.MenuItem) error {
	return r.baseRepository.Update(ctx, item)
}

// Delete soft deletes a menu item
func (r *menuItemRepository) Delete(ctx context.Context, item *model.MenuItem) error {
	return r.baseRepository.Delete(ctx, item)
}

// HardDelete permanently deletes a menu item
func (r *menuItemRepository) HardDelete(ctx context.Context, item *model.MenuItem) error {
	return r.baseRepository.HardDelete(ctx, item)
}
