package service

import (
	"context"
	"time"

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

	err := query.Find(&menus).Error
	return menus, err
}

// GetMenuByID retrieves a menu by its ID
func (s *menuService) GetMenuByID(ctx context.Context, id uint) (*model.MealMenu, error) {
	var menu model.MealMenu
	err := s.db.WithContext(ctx).
		Preload("CreatedByUser").
		Preload("MenuItems.MenuItem").
		First(&menu, id).Error

	if err != nil {
		return nil, err
	}

	return &menu, nil
}

// CreateMenu creates a new menu
func (s *menuService) CreateMenu(ctx context.Context, menu *model.MealMenu) error {
	return s.db.WithContext(ctx).Create(menu).Error
}

// UpdateMenu updates an existing menu
func (s *menuService) UpdateMenu(ctx context.Context, menu *model.MealMenu) error {
	return s.db.WithContext(ctx).Save(menu).Error
}

// DeleteMenu deletes a menu by its ID
func (s *menuService) DeleteMenu(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&model.MealMenu{}, id).Error
}

// GetMenuItems retrieves all menu items for a specific menu, optionally filtered by date
func (s *menuService) GetMenuItems(ctx context.Context, menuID uint, date *time.Time) ([]model.MenuItem, error) {
	var menu model.MealMenu
	err := s.db.WithContext(ctx).First(&menu, menuID).Error
	if err != nil {
		return nil, err
	}

	var menuItems []model.MenuItem
	for _, mealMenuItem := range menu.MenuItems {
		menuItems = append(menuItems, mealMenuItem.MenuItem)
	}

	return menuItems, nil
}

// AddMenuItem adds a menu item to a menu
func (s *menuService) AddMenuItem(ctx context.Context, menuID uint, menuItemID uint, setName string) error {
	mealMenuItem := model.MealMenuItem{
		MealMenuID: menuID,
		MenuItemID: menuItemID,
		SetName:    setName,
	}

	return s.db.WithContext(ctx).Create(&mealMenuItem).Error
}

// RemoveMenuItem removes a menu item from a menu
func (s *menuService) RemoveMenuItem(ctx context.Context, menuID uint, menuItemID uint) error {
	return s.db.WithContext(ctx).
		Where("meal_menu_id = ? AND menu_item_id = ?", menuID, menuItemID).
		Delete(&model.MealMenuItem{}).Error
}

func (s *menuService) CreateMenuItem(item *model.MenuItem) error {
	return s.db.Create(item).Error
}

func (s *menuService) GetMenuItemsByDate(date time.Time) ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := s.db.Where("date = ? AND is_active = ?", date, true).Find(&items).Error
	return items, err
}

func (s *menuService) UpdateMenuItem(id uint, item *model.MenuItem) error {
	return s.db.Model(&model.MenuItem{}).Where("id = ?", id).Updates(item).Error
}

func (s *menuService) DeleteMenuItem(id uint) error {
	return s.db.Model(&model.MenuItem{}).Where("id = ?", id).Update("is_active", false).Error
}

func (s *menuService) CreateMealRequest(request *model.MealRequest) error {
	return s.db.Create(request).Error
}

func (s *menuService) GetMealRequestsByDate(date time.Time) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := s.db.Preload("User").Preload("MenuItem").
		Where("date = ?", date).
		Find(&requests).Error
	return requests, err
}

func (s *menuService) GetMealRequestStats(date time.Time) (map[string]int, error) {
	var requests []model.MealRequest
	stats := make(map[string]int)

	err := s.db.Preload("User").
		Where("date = ?", date).
		Find(&requests).Error
	if err != nil {
		return nil, err
	}

	// Count by department
	for _, req := range requests {
		stats[req.User.Department]++
	}

	return stats, nil
}
