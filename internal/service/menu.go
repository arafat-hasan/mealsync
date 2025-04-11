package service

import (
	"context"
	"errors"
	"time"

	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// MenuService defines the interface for menu-related operations
type MenuService interface {
	GetMenus(ctx context.Context, date *time.Time) ([]model.MealMenu, error)
	GetMenuByID(ctx context.Context, id uint) (*model.MealMenu, error)
	CreateMenu(ctx context.Context, menu *model.MealMenu) error
	UpdateMenu(ctx context.Context, menu *model.MealMenu) error
	DeleteMenu(ctx context.Context, id uint) error
	GetMenuItems(ctx context.Context, menuID uint, date *time.Time) ([]model.MenuItem, error)
	AddMenuItem(ctx context.Context, menuID uint, menuItemID uint, setName string) error
	RemoveMenuItem(ctx context.Context, menuID uint, menuItemID uint) error
	CreateMenuItem(ctx context.Context, item *model.MenuItem) error
	GetMenuItemsByDate(ctx context.Context, date time.Time) ([]model.MenuItem, error)
	UpdateMenuItem(ctx context.Context, id uint, item *model.MenuItem) error
	DeleteMenuItem(ctx context.Context, id uint) error
	CreateMealRequest(ctx context.Context, request *model.MealRequest) error
	GetMealRequestsByDate(ctx context.Context, date time.Time) ([]model.MealRequest, error)
	GetMealRequestStats(ctx context.Context, date time.Time) (map[string]int, error)
}

// menuService implements the MenuService interface
type menuService struct {
	db *gorm.DB
}

// NewMenuService creates a new instance of MenuService
func NewMenuService(db *gorm.DB) MenuService {
	return &menuService{db: db}
}

// GetMenus retrieves all menus, optionally filtered by date
func (s *menuService) GetMenus(ctx context.Context, date *time.Time) ([]model.MealMenu, error) {
	var menus []model.MealMenu
	query := s.db.WithContext(ctx).Preload("CreatedByUser").Preload("MenuItems.MenuItem")

	if date != nil {
		query = query.Where("date = ?", date)
	}

	if err := query.Find(&menus).Error; err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch menus", err)
	}
	return menus, nil
}

// GetMenuByID retrieves a menu by its ID
func (s *menuService) GetMenuByID(ctx context.Context, id uint) (*model.MealMenu, error) {
	if id == 0 {
		return nil, apperrors.NewValidationError("Menu ID is required", nil)
	}

	var menu model.MealMenu
	err := s.db.WithContext(ctx).
		Preload("CreatedByUser").
		Preload("MenuItems.MenuItem").
		First(&menu, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Menu not found", err)
		}
		return nil, apperrors.NewInternalError("Failed to fetch menu", err)
	}

	return &menu, nil
}

// CreateMenu creates a new menu
func (s *menuService) CreateMenu(ctx context.Context, menu *model.MealMenu) error {
	if menu == nil {
		return apperrors.NewValidationError("Menu is required", nil)
	}

	if menu.Date.IsZero() {
		return apperrors.NewValidationError("Menu date is required", nil)
	}

	if err := s.db.WithContext(ctx).Create(menu).Error; err != nil {
		return apperrors.NewInternalError("Failed to create menu", err)
	}

	return nil
}

// UpdateMenu updates an existing menu
func (s *menuService) UpdateMenu(ctx context.Context, menu *model.MealMenu) error {
	if menu == nil {
		return apperrors.NewValidationError("Menu is required", nil)
	}

	if menu.ID == 0 {
		return apperrors.NewValidationError("Menu ID is required", nil)
	}

	// Check if menu exists
	if _, err := s.GetMenuByID(ctx, menu.ID); err != nil {
		return err
	}

	if err := s.db.WithContext(ctx).Save(menu).Error; err != nil {
		return apperrors.NewInternalError("Failed to update menu", err)
	}

	return nil
}

// DeleteMenu deletes a menu by its ID
func (s *menuService) DeleteMenu(ctx context.Context, id uint) error {
	if id == 0 {
		return apperrors.NewValidationError("Menu ID is required", nil)
	}

	// Check if menu exists
	if _, err := s.GetMenuByID(ctx, id); err != nil {
		return err
	}

	if err := s.db.WithContext(ctx).Delete(&model.MealMenu{}, id).Error; err != nil {
		return apperrors.NewInternalError("Failed to delete menu", err)
	}

	return nil
}

// GetMenuItems retrieves all menu items for a specific menu, optionally filtered by date
func (s *menuService) GetMenuItems(ctx context.Context, menuID uint, date *time.Time) ([]model.MenuItem, error) {
	if menuID == 0 {
		return nil, apperrors.NewValidationError("Menu ID is required", nil)
	}

	var menu model.MealMenu
	err := s.db.WithContext(ctx).First(&menu, menuID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Menu not found", err)
		}
		return nil, apperrors.NewInternalError("Failed to fetch menu", err)
	}

	var menuItems []model.MenuItem
	for _, mealMenuItem := range menu.MenuItems {
		menuItems = append(menuItems, mealMenuItem.MenuItem)
	}

	return menuItems, nil
}

// AddMenuItem adds a menu item to a menu
func (s *menuService) AddMenuItem(ctx context.Context, menuID uint, menuItemID uint, setName string) error {
	if menuID == 0 {
		return apperrors.NewValidationError("Menu ID is required", nil)
	}

	if menuItemID == 0 {
		return apperrors.NewValidationError("Menu item ID is required", nil)
	}

	// Check if menu exists
	if _, err := s.GetMenuByID(ctx, menuID); err != nil {
		return err
	}

	mealMenuItem := model.MealMenuItem{
		MealMenuID: menuID,
		MenuItemID: menuItemID,
		SetName:    setName,
	}

	if err := s.db.WithContext(ctx).Create(&mealMenuItem).Error; err != nil {
		return apperrors.NewInternalError("Failed to add menu item", err)
	}

	return nil
}

// RemoveMenuItem removes a menu item from a menu
func (s *menuService) RemoveMenuItem(ctx context.Context, menuID uint, menuItemID uint) error {
	if menuID == 0 {
		return apperrors.NewValidationError("Menu ID is required", nil)
	}

	if menuItemID == 0 {
		return apperrors.NewValidationError("Menu item ID is required", nil)
	}

	// Check if menu exists
	if _, err := s.GetMenuByID(ctx, menuID); err != nil {
		return err
	}

	if err := s.db.WithContext(ctx).
		Where("meal_menu_id = ? AND menu_item_id = ?", menuID, menuItemID).
		Delete(&model.MealMenuItem{}).Error; err != nil {
		return apperrors.NewInternalError("Failed to remove menu item", err)
	}

	return nil
}

// CreateMenuItem creates a new menu item
func (s *menuService) CreateMenuItem(ctx context.Context, item *model.MenuItem) error {
	if item == nil {
		return apperrors.NewValidationError("Menu item is required", nil)
	}

	if item.Name == "" {
		return apperrors.NewValidationError("Menu item name is required", nil)
	}

	if err := s.db.WithContext(ctx).Create(item).Error; err != nil {
		return apperrors.NewInternalError("Failed to create menu item", err)
	}

	return nil
}

// GetMenuItemsByDate retrieves menu items by date
func (s *menuService) GetMenuItemsByDate(ctx context.Context, date time.Time) ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := s.db.WithContext(ctx).Where("date = ? AND is_active = ?", date, true).Find(&items).Error
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch menu items", err)
	}
	return items, nil
}

// UpdateMenuItem updates a menu item
func (s *menuService) UpdateMenuItem(ctx context.Context, id uint, item *model.MenuItem) error {
	if id == 0 {
		return apperrors.NewValidationError("Menu item ID is required", nil)
	}

	if item == nil {
		return apperrors.NewValidationError("Menu item is required", nil)
	}

	if err := s.db.WithContext(ctx).Model(&model.MenuItem{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return apperrors.NewInternalError("Failed to update menu item", err)
	}

	return nil
}

// DeleteMenuItem deletes a menu item
func (s *menuService) DeleteMenuItem(ctx context.Context, id uint) error {
	if id == 0 {
		return apperrors.NewValidationError("Menu item ID is required", nil)
	}

	if err := s.db.WithContext(ctx).Model(&model.MenuItem{}).Where("id = ?", id).Update("is_active", false).Error; err != nil {
		return apperrors.NewInternalError("Failed to delete menu item", err)
	}

	return nil
}

// CreateMealRequest creates a new meal request
func (s *menuService) CreateMealRequest(ctx context.Context, request *model.MealRequest) error {
	if request == nil {
		return apperrors.NewValidationError("Meal request is required", nil)
	}

	if request.UserID == 0 {
		return apperrors.NewValidationError("User ID is required", nil)
	}

	if request.MenuItemID == 0 {
		return apperrors.NewValidationError("Menu item ID is required", nil)
	}

	if request.RequestedFor.IsZero() {
		return apperrors.NewValidationError("Request date is required", nil)
	}

	if err := s.db.WithContext(ctx).Create(request).Error; err != nil {
		return apperrors.NewInternalError("Failed to create meal request", err)
	}

	return nil
}

// GetMealRequestsByDate retrieves meal requests by date
func (s *menuService) GetMealRequestsByDate(ctx context.Context, date time.Time) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := s.db.WithContext(ctx).Preload("User").Preload("MenuItem").
		Where("date = ?", date).
		Find(&requests).Error
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch meal requests", err)
	}
	return requests, nil
}

// GetMealRequestStats retrieves meal request statistics by date
func (s *menuService) GetMealRequestStats(ctx context.Context, date time.Time) (map[string]int, error) {
	var requests []model.MealRequest
	stats := make(map[string]int)

	err := s.db.WithContext(ctx).Preload("User").
		Where("date = ?", date).
		Find(&requests).Error
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch meal request statistics", err)
	}

	// Count by department
	for _, req := range requests {
		stats[req.User.Department]++
	}

	return stats, nil
}
