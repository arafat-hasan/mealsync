package service

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// mealEventService handles business logic for meal event operations
type mealEventService struct {
	mealRepo     repository.MealEventRepository
	userRepo     repository.UserRepository
	menuRepo     repository.MenuSetRepository
	addressRepo  repository.EventAddressRepository
	requestRepo  repository.MealRequestRepository
	commentRepo  repository.MealCommentRepository
	notifService NotificationService
}

// NewMealEventService creates a new instance of MealEventService
func NewMealEventService(
	mealRepo repository.MealEventRepository,
	userRepo repository.UserRepository,
	menuRepo repository.MenuSetRepository,
	addressRepo repository.EventAddressRepository,
	requestRepo repository.MealRequestRepository,
	commentRepo repository.MealCommentRepository,
	notifService NotificationService,
) MealEventService {
	return &mealEventService{
		mealRepo:     mealRepo,
		userRepo:     userRepo,
		menuRepo:     menuRepo,
		addressRepo:  addressRepo,
		requestRepo:  requestRepo,
		commentRepo:  commentRepo,
		notifService: notifService,
	}
}

// GetMeals retrieves meal events based on user role and filters
func (s *mealEventService) GetMeals(ctx context.Context, userID uint, isAdmin bool) ([]model.MealEvent, error) {
	if isAdmin {
		return s.mealRepo.FindAll(ctx)
	}
	return s.mealRepo.FindByUserID(ctx, userID)
}

// GetMealByID retrieves a specific meal event by ID
func (s *mealEventService) GetMealByID(ctx context.Context, id uint, userID uint, isAdmin bool) (*model.MealEvent, error) {
	meal, err := s.mealRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !isAdmin && meal.CreatedBy != userID {
		return nil, errors.NewForbiddenError("unauthorized to access this meal event", nil)
	}

	return meal, nil
}

// CreateMeal creates a new meal event
func (s *mealEventService) CreateMeal(ctx context.Context, meal *model.MealEvent, userID uint) error {
	if meal == nil {
		return errors.NewValidationError("meal event cannot be nil", nil)
	}

	if meal.Name == "" {
		return errors.NewValidationError("meal event name is required", nil)
	}

	if meal.EventDate.IsZero() {
		return errors.NewValidationError("meal event date is required", nil)
	}

	if meal.CutoffTime.IsZero() {
		return errors.NewValidationError("cutoff time is required", nil)
	}

	// Set created by
	meal.CreatedBy = userID
	meal.UpdatedBy = userID

	return s.mealRepo.Create(ctx, meal)
}

// UpdateMeal updates an existing meal event
func (s *mealEventService) UpdateMeal(ctx context.Context, id uint, meal *model.MealEvent, userID uint) error {
	if meal == nil {
		return errors.NewValidationError("meal event cannot be nil", nil)
	}

	existingMeal, err := s.mealRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership
	if existingMeal.CreatedBy != userID {
		return errors.NewForbiddenError("unauthorized to update this meal event", nil)
	}

	// Update fields
	existingMeal.Name = meal.Name
	existingMeal.Description = meal.Description
	existingMeal.EventDate = meal.EventDate
	existingMeal.CutoffTime = meal.CutoffTime
	existingMeal.UpdatedBy = userID

	return s.mealRepo.Update(ctx, existingMeal)
}

// DeleteMeal soft deletes a meal event
func (s *mealEventService) DeleteMeal(ctx context.Context, id uint, userID uint) error {
	meal, err := s.mealRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership
	if meal.CreatedBy != userID {
		return errors.NewForbiddenError("unauthorized to delete this meal event", nil)
	}

	meal.UpdatedBy = userID
	return s.mealRepo.Delete(ctx, meal)
}

// SubmitMealRequest submits a new meal request
func (s *mealEventService) SubmitMealRequest(ctx context.Context, request *model.MealRequest, userID uint) error {
	if request == nil {
		return errors.NewValidationError("request cannot be nil", nil)
	}

	// Verify meal event exists and is active
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

	// Set request fields
	request.UserID = userID
	request.CreatedBy = userID
	request.UpdatedBy = userID

	return s.requestRepo.Create(ctx, request)
}

// WithdrawMealRequest withdraws a meal request
func (s *mealEventService) WithdrawMealRequest(ctx context.Context, requestID uint, userID uint) error {
	request, err := s.requestRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	// Verify ownership
	if request.UserID != userID {
		return errors.NewForbiddenError("unauthorized to withdraw this request", nil)
	}

	request.UpdatedBy = userID
	return s.requestRepo.Delete(ctx, request)
}

// GetMealRequests retrieves meal requests based on user role and filters
func (s *mealEventService) GetMealRequests(ctx context.Context, userID uint, isAdmin bool) ([]model.MealRequest, error) {
	if isAdmin {
		return s.requestRepo.FindAll(ctx)
	}
	return s.requestRepo.FindByUserID(ctx, userID)
}
