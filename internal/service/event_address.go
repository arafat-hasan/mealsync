package service

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// eventAddressService handles business logic for event address operations
type eventAddressService struct {
	addressRepo repository.EventAddressRepository
	userRepo    repository.UserRepository
}

// NewEventAddressService creates a new instance of EventAddressService
func NewEventAddressService(
	addressRepo repository.EventAddressRepository,
	userRepo repository.UserRepository,
) EventAddressService {
	return &eventAddressService{
		addressRepo: addressRepo,
		userRepo:    userRepo,
	}
}

// GetAddresses retrieves all event addresses
func (s *eventAddressService) GetAddresses(ctx context.Context) ([]model.MealEventAddress, error) {
	return s.addressRepo.FindActive(ctx, map[string]interface{}{
		"is_active": true,
	})
}

// GetAddressByID retrieves a specific event address by ID
func (s *eventAddressService) GetAddressByID(ctx context.Context, id uint) (*model.MealEventAddress, error) {
	address, err := s.addressRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("event address not found", err)
	}
	return address, nil
}

// CreateAddress creates a new event address
func (s *eventAddressService) CreateAddress(ctx context.Context, address *model.MealEventAddress, userID uint) error {
	if address == nil {
		return errors.NewValidationError("address cannot be nil", nil)
	}

	if err := s.validateAddress(address); err != nil {
		return err
	}

	// Set metadata
	address.CreatedBy = userID
	address.UpdatedBy = userID

	if err := s.addressRepo.Create(ctx, address); err != nil {
		return errors.NewInternalError("failed to create event address", err)
	}

	return nil
}

// UpdateAddress updates an existing event address
func (s *eventAddressService) UpdateAddress(ctx context.Context, id uint, address *model.MealEventAddress, userID uint) error {
	if address == nil {
		return errors.NewValidationError("address cannot be nil", nil)
	}

	existingAddress, err := s.addressRepo.FindByID(ctx, id)
	if err != nil {
		return errors.NewNotFoundError("event address not found", err)
	}

	if err := s.validateAddress(address); err != nil {
		return err
	}

	// Update fields
	existingAddress.MealEventID = address.MealEventID
	existingAddress.AddressID = address.AddressID
	existingAddress.UpdatedBy = userID

	if err := s.addressRepo.Update(ctx, existingAddress); err != nil {
		return errors.NewInternalError("failed to update event address", err)
	}

	return nil
}

// DeleteAddress soft deletes an event address
func (s *eventAddressService) DeleteAddress(ctx context.Context, id uint, userID uint) error {
	address, err := s.addressRepo.FindByID(ctx, id)
	if err != nil {
		return errors.NewNotFoundError("event address not found", err)
	}

	address.UpdatedBy = userID
	if err := s.addressRepo.Delete(ctx, address); err != nil {
		return errors.NewInternalError("failed to delete event address", err)
	}

	return nil
}

// GetAddressesByType retrieves event addresses by type
func (s *eventAddressService) GetAddressesByType(ctx context.Context, addressType string) ([]model.MealEventAddress, error) {
	if addressType == "" {
		return nil, errors.NewValidationError("address type is required", nil)
	}

	addresses, err := s.addressRepo.FindByAddressType(ctx, addressType)
	if err != nil {
		return nil, errors.NewInternalError("failed to fetch addresses by type", err)
	}

	return addresses, nil
}

// GetAddressesByCapacity retrieves event addresses by capacity range
func (s *eventAddressService) GetAddressesByCapacity(ctx context.Context, minCapacity, maxCapacity int) ([]model.MealEventAddress, error) {
	if minCapacity < 0 || maxCapacity < minCapacity {
		return nil, errors.NewValidationError("invalid capacity range", nil)
	}

	addresses, err := s.addressRepo.FindByCapacity(ctx, minCapacity, maxCapacity)
	if err != nil {
		return nil, errors.NewInternalError("failed to fetch addresses by capacity", err)
	}

	return addresses, nil
}

// GetAddressesByLocation retrieves event addresses within a radius of given coordinates
func (s *eventAddressService) GetAddressesByLocation(ctx context.Context, latitude, longitude, radius float64) ([]model.MealEventAddress, error) {
	if radius <= 0 {
		return nil, errors.NewValidationError("radius must be positive", nil)
	}

	addresses, err := s.addressRepo.FindByLocation(ctx, latitude, longitude, radius)
	if err != nil {
		return nil, errors.NewInternalError("failed to fetch addresses by location", err)
	}

	return addresses, nil
}

// GetAvailableAddresses retrieves available addresses for a given date
func (s *eventAddressService) GetAvailableAddresses(ctx context.Context, date time.Time) ([]model.MealEventAddress, error) {
	if date.IsZero() {
		return nil, errors.NewValidationError("date is required", nil)
	}

	addresses, err := s.addressRepo.FindAvailableAddresses(ctx, date)
	if err != nil {
		return nil, errors.NewInternalError("failed to fetch available addresses", err)
	}

	return addresses, nil
}

// validateAddress performs validation on the event address
func (s *eventAddressService) validateAddress(address *model.MealEventAddress) error {
	if address.MealEventID == 0 {
		return errors.NewValidationError("meal event ID is required", nil)
	}

	if address.AddressID == 0 {
		return errors.NewValidationError("address ID is required", nil)
	}

	return nil
}
