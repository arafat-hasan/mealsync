package service

import (
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{db: db}
}

func (s *MenuService) CreateMenuItem(item *model.MenuItem) error {
	return s.db.Create(item).Error
}

func (s *MenuService) GetMenuItemsByDate(date time.Time) ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := s.db.Where("date = ? AND is_active = ?", date, true).Find(&items).Error
	return items, err
}

func (s *MenuService) UpdateMenuItem(id uint, item *model.MenuItem) error {
	return s.db.Model(&model.MenuItem{}).Where("id = ?", id).Updates(item).Error
}

func (s *MenuService) DeleteMenuItem(id uint) error {
	return s.db.Model(&model.MenuItem{}).Where("id = ?", id).Update("is_active", false).Error
}

func (s *MenuService) CreateMealRequest(request *model.MealRequest) error {
	return s.db.Create(request).Error
}

func (s *MenuService) GetMealRequestsByDate(date time.Time) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := s.db.Preload("User").Preload("MenuItem").
		Where("date = ?", date).
		Find(&requests).Error
	return requests, err
}

func (s *MenuService) GetMealRequestStats(date time.Time) (map[string]int, error) {
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
