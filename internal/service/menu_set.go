package service

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// menuSetService handles business logic for menu set and menu item operations
type menuSetService struct {
	menuRepo     repository.MenuSetRepository
	menuItemRepo repository.MenuItemRepository
	userRepo     repository.UserRepository
}

// NewMenuSetService creates a new instance of MenuSetService
func NewMenuSetService(
	menuRepo repository.MenuSetRepository,
	menuItemRepo repository.MenuItemRepository,
	userRepo repository.UserRepository,
) MenuSetService {
	return &menuSetService{
		menuRepo:     menuRepo,
		menuItemRepo: menuItemRepo,
		userRepo:     userRepo,
	}
}

// GetMenuSets retrieves all menu sets
func (s *menuSetService) GetMenuSets(ctx context.Context) ([]model.MenuSet, error) {
	return s.menuRepo.FindActive(ctx, map[string]interface{}{
		"is_active": true,
	})
}

// GetMenuSetByID retrieves a specific menu set by ID
func (s *menuSetService) GetMenuSetByID(ctx context.Context, id uint) (*model.MenuSet, error) {
	return s.menuRepo.FindByID(ctx, id)
}

// CreateMenuSet creates a new menu set
func (s *menuSetService) CreateMenuSet(ctx context.Context, menuSet *model.MenuSet, userID uint) error {
	if menuSet == nil {
		return errors.NewValidationError("menu set cannot be nil", nil)
	}

	if menuSet.MenuSetName == "" {
		return errors.NewValidationError("menu set name is required", nil)
	}

	// Set created by
	menuSet.CreatedBy = userID
	menuSet.UpdatedBy = userID

	return s.menuRepo.Create(ctx, menuSet)
}

// UpdateMenuSet updates an existing menu set
func (s *menuSetService) UpdateMenuSet(ctx context.Context, id uint, menuSet *model.MenuSet, userID uint) error {
	if menuSet == nil {
		return errors.NewValidationError("menu set cannot be nil", nil)
	}

	existingMenuSet, err := s.menuRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Update fields
	existingMenuSet.MenuSetName = menuSet.MenuSetName
	existingMenuSet.MenuSetDescription = menuSet.MenuSetDescription
	existingMenuSet.UpdatedBy = userID

	return s.menuRepo.Update(ctx, existingMenuSet)
}

// DeleteMenuSet soft deletes a menu set
func (s *menuSetService) DeleteMenuSet(ctx context.Context, id uint, userID uint) error {
	menuSet, err := s.menuRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	menuSet.UpdatedBy = userID
	return s.menuRepo.Delete(ctx, menuSet)
}

// GetMenuItems retrieves all menu items
func (s *menuSetService) GetMenuItems(ctx context.Context) ([]model.MenuItem, error) {
	return s.menuItemRepo.FindActive(ctx, map[string]interface{}{
		"is_active": true,
	})
}

// GetMenuItemByID retrieves a specific menu item by ID
func (s *menuSetService) GetMenuItemByID(ctx context.Context, id uint) (*model.MenuItem, error) {
	return s.menuItemRepo.FindByID(ctx, id)
}

// CreateMenuItem creates a new menu item
func (s *menuSetService) CreateMenuItem(ctx context.Context, menuItem *model.MenuItem, userID uint) error {
	if menuItem == nil {
		return errors.NewValidationError("menu item cannot be nil", nil)
	}

	if menuItem.Name == "" {
		return errors.NewValidationError("menu item name is required", nil)
	}

	// Set created by
	menuItem.CreatedBy = userID
	menuItem.UpdatedBy = userID

	return s.menuItemRepo.Create(ctx, menuItem)
}

// UpdateMenuItem updates an existing menu item
func (s *menuSetService) UpdateMenuItem(ctx context.Context, id uint, menuItem *model.MenuItem, userID uint) error {
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
func (s *menuSetService) DeleteMenuItem(ctx context.Context, id uint, userID uint) error {
	menuItem, err := s.menuItemRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	menuItem.UpdatedBy = userID
	return s.menuItemRepo.Delete(ctx, menuItem)
}

// AddItemToMenuSet adds a menu item to a menu set
func (s *menuSetService) AddItemToMenuSet(ctx context.Context, menuSetID uint, menuItemID uint, userID uint) error {
	// Verify menu set exists
	if _, err := s.menuRepo.FindByID(ctx, menuSetID); err != nil {
		return err
	}

	// Verify menu item exists
	if _, err := s.menuItemRepo.FindByID(ctx, menuItemID); err != nil {
		return err
	}

	// Create menu set item
	menuSetItem := &model.MenuSetItem{
		MenuSetID:  menuSetID,
		MenuItemID: menuItemID,
		CreatedBy:  userID,
		UpdatedBy:  userID,
	}

	return s.menuRepo.AddMenuItem(ctx, menuSetItem)
}

// RemoveItemFromMenuSet removes a menu item from a menu set
func (s *menuSetService) RemoveItemFromMenuSet(ctx context.Context, menuSetID uint, menuItemID uint, userID uint) error {
	menuSetItem := &model.MenuSetItem{
		MenuSetID:  menuSetID,
		MenuItemID: menuItemID,
		UpdatedBy:  userID,
	}

	return s.menuRepo.RemoveMenuItem(ctx, menuSetItem)
}

// GetMenuSetItems retrieves all menu items in a menu set
func (s *menuSetService) GetMenuSetItems(ctx context.Context, menuSetID uint) ([]model.MenuItem, error) {
	return s.menuRepo.FindMenuItems(ctx, menuSetID)
}
