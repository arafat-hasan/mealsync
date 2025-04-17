package repository

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// eventAddressRepository implements EventAddressRepository interface
type eventAddressRepository struct {
	*baseRepository[model.MealEventAddress]
	db *gorm.DB
}

// NewEventAddressRepository creates a new instance of EventAddressRepository
func NewEventAddressRepository(db *gorm.DB) EventAddressRepository {
	return &eventAddressRepository{
		baseRepository: NewBaseRepository[model.MealEventAddress](db),
		db:             db,
	}
}

// Create creates a new event address
func (r *eventAddressRepository) Create(ctx context.Context, address *model.MealEventAddress) error {
	return r.baseRepository.Create(ctx, address)
}

// FindByID finds an event address by ID
func (r *eventAddressRepository) FindByID(ctx context.Context, id uint) (*model.MealEventAddress, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all event addresses
func (r *eventAddressRepository) FindAll(ctx context.Context) ([]model.MealEventAddress, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active event addresses based on conditions
func (r *eventAddressRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealEventAddress, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates an event address
func (r *eventAddressRepository) Update(ctx context.Context, address *model.MealEventAddress) error {
	return r.baseRepository.Update(ctx, address)
}

// Delete soft deletes an event address
func (r *eventAddressRepository) Delete(ctx context.Context, address *model.MealEventAddress) error {
	return r.baseRepository.Delete(ctx, address)
}

// HardDelete permanently deletes an event address
func (r *eventAddressRepository) HardDelete(ctx context.Context, address *model.MealEventAddress) error {
	return r.baseRepository.HardDelete(ctx, address)
}

// FindByMealEventID finds event addresses by meal event ID
func (r *eventAddressRepository) FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	err := r.db.WithContext(ctx).Where("meal_event_id = ?", mealEventID).Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// CountByMealEventID counts event addresses by meal event ID
func (r *eventAddressRepository) CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.MealEventAddress{}).Where("meal_event_id = ?", mealEventID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindAvailableAddresses finds available addresses for a given date
func (r *eventAddressRepository) FindAvailableAddresses(ctx context.Context, date time.Time) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	err := r.db.WithContext(ctx).
		Where("is_available = ? AND DATE(created_at) = DATE(?)", true, date).
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindByAddressType finds addresses by type
func (r *eventAddressRepository) FindByAddressType(ctx context.Context, addressType string) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	err := r.db.WithContext(ctx).
		Where("address_type = ?", addressType).
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindByBuildingName finds addresses by building name
func (r *eventAddressRepository) FindByBuildingName(ctx context.Context, buildingName string) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	err := r.db.WithContext(ctx).
		Where("building_name = ?", buildingName).
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindByCapacity finds addresses by capacity range
func (r *eventAddressRepository) FindByCapacity(ctx context.Context, minCapacity, maxCapacity int) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	err := r.db.WithContext(ctx).
		Where("capacity >= ? AND capacity <= ?", minCapacity, maxCapacity).
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindByLocation finds addresses within a radius of given coordinates
func (r *eventAddressRepository) FindByLocation(ctx context.Context, lat, lng, radius float64) ([]model.MealEventAddress, error) {
	var addresses []model.MealEventAddress
	// Using the Haversine formula to find addresses within the radius
	query := `
		SELECT * FROM meal_event_addresses 
		WHERE (6371 * acos(cos(radians(?)) * cos(radians(latitude)) * 
		cos(radians(longitude) - radians(?)) + sin(radians(?)) * 
		sin(radians(latitude)))) <= ?`
	err := r.db.WithContext(ctx).
		Raw(query, lat, lng, lat, radius).
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindWithEventDetails finds an address with its associated event details by meal event ID
func (r *eventAddressRepository) FindWithEventDetails(ctx context.Context, mealEventID uint) (*model.MealEventAddress, error) {
	var address model.MealEventAddress
	err := r.db.WithContext(ctx).
		Preload("MealEvent").
		Where("meal_event_id = ?", mealEventID).
		First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}
