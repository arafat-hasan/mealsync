package service

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// menuItemService handles business logic for menu item operations
type menuItemService struct {
	menuItemRepo repository.MenuItemRepository
	userRepo     repository.UserRepository
}

// NewMenuItemService creates a new instance of MenuItemService
func NewMenuItemService(
	menuItemRepo repository.MenuItemRepository,
	userRepo repository.UserRepository,
) MenuItemService {
	return &menuItemService{
		menuItemRepo: menuItemRepo,
		userRepo:     userRepo,
	}
}

// GetMenuItems retrieves all menu items
func (s *menuItemService) GetMenuItems(ctx context.Context) ([]model.MenuItem, error) {
	return s.menuItemRepo.FindActive(ctx, map[string]interface{}{
		"is_active": true,
	})
}

// GetMenuItemByID retrieves a specific menu item by ID
func (s *menuItemService) GetMenuItemByID(ctx context.Context, id uint) (*model.MenuItem, error) {
	return s.menuItemRepo.FindByID(ctx, id)
}

// CreateMenuItem creates a new menu item
func (s *menuItemService) CreateMenuItem(ctx context.Context, menuItem *model.MenuItem, userID uint) error {
	if menuItem == nil {
		return errors.NewValidationError("menu item cannot be nil", nil)
	}

	if menuItem.Name == "" {
		return errors.NewValidationError("name is required", nil)
	}

	if menuItem.Description == "" {
		return errors.NewValidationError("description is required", nil)
	}

	// Set created by
	menuItem.CreatedBy = userID
	menuItem.UpdatedBy = userID

	return s.menuItemRepo.Create(ctx, menuItem)
}

// UpdateMenuItem updates an existing menu item
func (s *menuItemService) UpdateMenuItem(ctx context.Context, id uint, menuItem *model.MenuItem, userID uint) error {
	if menuItem == nil {
		return errors.NewValidationError("menu item cannot be nil", nil)
	}

	existingMenuItem, err := s.menuItemRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Update fields
	existingMenuItem.Name = menuItem.Name
	existingMenuItem.Description = menuItem.Description
	existingMenuItem.ImageURL = menuItem.ImageURL
	existingMenuItem.UpdatedBy = userID

	return s.menuItemRepo.Update(ctx, existingMenuItem)
}

// DeleteMenuItem soft deletes a menu item
func (s *menuItemService) DeleteMenuItem(ctx context.Context, id uint, userID uint) error {
	menuItem, err := s.menuItemRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	menuItem.UpdatedBy = userID
	return s.menuItemRepo.Delete(ctx, menuItem)
}

// GetMenuItemsByCategory retrieves menu items by category
func (s *menuItemService) GetMenuItemsByCategory(ctx context.Context, category string) ([]model.MenuItem, error) {
	return s.menuItemRepo.FindActive(ctx, map[string]interface{}{
		"is_active": true,
	})
}

// GetMenuItemsByMenuSet retrieves menu items for a specific menu set
func (s *menuItemService) GetMenuItemsByMenuSet(ctx context.Context, menuSetID uint) ([]model.MenuItem, error) {
	return s.menuItemRepo.FindActive(ctx, map[string]interface{}{
		"menu_set_id": menuSetID,
		"is_active":   true,
	})
}
