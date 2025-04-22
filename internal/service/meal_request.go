package service

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// mealRequestService handles business logic for meal request operations
type mealRequestService struct {
	requestRepo repository.MealRequestRepository
	mealRepo    repository.MealEventRepository
	userRepo    repository.UserRepository
}

// NewMealRequestService creates a new instance of MealRequestService
func NewMealRequestService(
	requestRepo repository.MealRequestRepository,
	mealRepo repository.MealEventRepository,
	userRepo repository.UserRepository,
) MealRequestService {
	return &mealRequestService{
		requestRepo: requestRepo,
		mealRepo:    mealRepo,
		userRepo:    userRepo,
	}
}

// GetMealRequests retrieves meal requests based on user role and filters
func (s *mealRequestService) GetMealRequests(ctx context.Context, userID uint, isAdmin bool) ([]model.MealRequest, error) {
	if isAdmin {
		return s.requestRepo.FindAll(ctx)
	}
	return s.requestRepo.FindByUserID(ctx, userID)
}

// GetMealRequestByID retrieves a specific meal request by ID
func (s *mealRequestService) GetMealRequestByID(ctx context.Context, id uint, userID uint, isAdmin bool) (*model.MealRequest, error) {
	request, err := s.requestRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !isAdmin && request.UserID != userID {
		return nil, errors.NewForbiddenError("unauthorized to access this request", nil)
	}

	return request, nil
}

// CreateMealRequest creates a new meal request
func (s *mealRequestService) CreateMealRequest(ctx context.Context, request *model.MealRequest, userID uint) error {
	if request == nil {
		return errors.NewValidationError("request cannot be nil", nil)
	}

	// Validate meal event exists and is active
	meal, err := s.mealRepo.FindByID(ctx, request.MealEventID)
	if err != nil {
		return err
	}

	if !meal.IsActive {
		return errors.NewValidationError("meal event is not active", nil)
	}

	// Check if cutoff time has passed
	if time.Now().After(meal.CutoffTime) {
		return errors.NewValidationError("cutoff time has passed", nil)
	}

	// Check if user already has a request for this meal event
	existingRequests, err := s.requestRepo.FindByMealEventID(ctx, request.MealEventID)
	if err != nil {
		return err
	}

	for _, existingRequest := range existingRequests {
		if existingRequest.UserID == userID {
			return errors.NewValidationError("user already has a request for this meal event", nil)
		}
	}

	// Set request fields
	request.UserID = userID
	request.CreatedBy = userID
	request.UpdatedBy = userID

	return s.requestRepo.Create(ctx, request)
}

// UpdateMealRequest updates an existing meal request
func (s *mealRequestService) UpdateMealRequest(ctx context.Context, id uint, request *model.MealRequest, userID uint, isAdmin bool) error {
	if request == nil {
		return errors.NewValidationError("request cannot be nil", nil)
	}

	existingRequest, err := s.requestRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership or admin status
	if !isAdmin && existingRequest.UserID != userID {
		return errors.NewForbiddenError("unauthorized to update this request", nil)
	}

	// Validate meal event exists and is active
	meal, err := s.mealRepo.FindByID(ctx, existingRequest.MealEventID)
	if err != nil {
		return err
	}

	if !meal.IsActive {
		return errors.NewValidationError("meal event is not active", nil)
	}

	// Check if cutoff time has passed
	if time.Now().After(meal.CutoffTime) {
		return errors.NewValidationError("cutoff time has passed", nil)
	}

	// Update fields
	existingRequest.MenuSetID = request.MenuSetID
	existingRequest.EventAddressID = request.EventAddressID
	existingRequest.UpdatedBy = userID

	return s.requestRepo.Update(ctx, existingRequest)
}

// DeleteMealRequest soft deletes a meal request
func (s *mealRequestService) DeleteMealRequest(ctx context.Context, id uint, userID uint, isAdmin bool) error {
	request, err := s.requestRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership or admin status
	if !isAdmin && request.UserID != userID {
		return errors.NewForbiddenError("unauthorized to delete this request", nil)
	}

	// Check if cutoff time has passed
	meal, err := s.mealRepo.FindByID(ctx, request.MealEventID)
	if err != nil {
		return err
	}

	if time.Now().After(meal.CutoffTime) {
		return errors.NewValidationError("cutoff time has passed", nil)
	}

	request.UpdatedBy = userID
	return s.requestRepo.Delete(ctx, request)
}

// AddRequestItem adds an item to a meal request
func (s *mealRequestService) AddRequestItem(ctx context.Context, requestID uint, item *model.MealRequestItem, userID uint, isAdmin bool) error {
	if item == nil {
		return errors.NewValidationError("item cannot be nil", nil)
	}

	request, err := s.requestRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	// Verify ownership or admin status
	if !isAdmin && request.UserID != userID {
		return errors.NewForbiddenError("unauthorized to modify this request", nil)
	}

	// Set item fields
	item.MealRequestID = requestID
	item.CreatedBy = userID
	item.UpdatedBy = userID

	return s.requestRepo.AddRequestItem(ctx, item)
}

// RemoveRequestItem removes an item from a meal request
func (s *mealRequestService) RemoveRequestItem(ctx context.Context, requestID uint, itemID uint, userID uint, isAdmin bool) error {
	request, err := s.requestRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	// Verify ownership or admin status
	if !isAdmin && request.UserID != userID {
		return errors.NewForbiddenError("unauthorized to modify this request", nil)
	}

	items, err := s.requestRepo.FindRequestItems(ctx, requestID)
	if err != nil {
		return err
	}

	var itemToRemove *model.MealRequestItem
	for _, item := range items {
		if item.ID == itemID {
			itemToRemove = &item
			break
		}
	}

	if itemToRemove == nil {
		return errors.NewNotFoundError("request item not found", nil)
	}

	itemToRemove.UpdatedBy = userID
	return s.requestRepo.RemoveRequestItem(ctx, itemToRemove)
}

// GetRequestItems retrieves all items for a meal request
func (s *mealRequestService) GetRequestItems(ctx context.Context, requestID uint, userID uint, isAdmin bool) ([]model.MealRequestItem, error) {
	request, err := s.requestRepo.FindByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	// Verify ownership or admin status
	if !isAdmin && request.UserID != userID {
		return nil, errors.NewForbiddenError("unauthorized to view this request", nil)
	}

	return s.requestRepo.FindRequestItems(ctx, requestID)
}

// UpdateRequestStatus updates the status of a meal request
func (s *mealRequestService) UpdateRequestStatus(ctx context.Context, requestID uint, status model.RequestStatus, userID uint) error {
	// Verify user exists and is admin
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.Role != model.UserRoleAdmin {
		return errors.NewForbiddenError("unauthorized to update request status", nil)
	}

	return s.requestRepo.UpdateRequestStatus(ctx, requestID, status)
}
