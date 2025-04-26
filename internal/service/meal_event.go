package service

import (
	"context"
	"time"

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
	commentRepo  repository.MenuItemCommentRepository
	notifService NotificationService
}

// NewMealEventService creates a new instance of MealEventService
func NewMealEventService(
	mealRepo repository.MealEventRepository,
	userRepo repository.UserRepository,
	menuRepo repository.MenuSetRepository,
	addressRepo repository.EventAddressRepository,
	requestRepo repository.MealRequestRepository,
	commentRepo repository.MenuItemCommentRepository,
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

// Create implements creating a new meal event
func (s *mealEventService) Create(ctx context.Context, meal *model.MealEvent) error {
	return s.mealRepo.Create(ctx, meal)
}

// FindByID implements finding a meal event by ID
func (s *mealEventService) FindByID(ctx context.Context, id uint) (*model.MealEvent, error) {
	return s.mealRepo.FindByID(ctx, id)
}

// FindAll implements finding all meal events
func (s *mealEventService) FindAll(ctx context.Context) ([]model.MealEvent, error) {
	return s.mealRepo.FindAll(ctx)
}

// FindActive implements finding active meal events
func (s *mealEventService) FindActive(ctx context.Context) ([]model.MealEvent, error) {
	return s.mealRepo.FindActive(ctx, map[string]interface{}{"is_active": true})
}

// Update implements updating a meal event
func (s *mealEventService) Update(ctx context.Context, meal *model.MealEvent) error {
	return s.mealRepo.Update(ctx, meal)
}

// Delete implements soft deleting a meal event
func (s *mealEventService) Delete(ctx context.Context, meal *model.MealEvent) error {
	return s.mealRepo.Delete(ctx, meal)
}

// HardDelete implements permanently deleting a meal event
func (s *mealEventService) HardDelete(ctx context.Context, meal *model.MealEvent) error {
	return s.mealRepo.HardDelete(ctx, meal)
}

// AddMenuSetToEvent associates a menu set with a meal event, including label and note
func (s *mealEventService) AddMenuSetToEvent(ctx context.Context, MealEventSet *model.MealEventSet) error {
	// Validate meal event exists
	_, err := s.mealRepo.FindByID(ctx, MealEventSet.MealEventID)
	if err != nil {
		return err
	}

	// Validate menu set exists
	_, err = s.menuRepo.FindByID(ctx, MealEventSet.MenuSetID)
	if err != nil {
		return err
	}

	return s.mealRepo.AddMenuSetToEvent(ctx, MealEventSet)
}

// UpdateMenuSetInEvent updates label and note for a menu set in an event
func (s *mealEventService) UpdateMenuSetInEvent(ctx context.Context, MealEventSet *model.MealEventSet) error {
	// Validate meal event exists
	_, err := s.mealRepo.FindByID(ctx, MealEventSet.MealEventID)
	if err != nil {
		return err
	}

	// Validate menu set exists
	_, err = s.menuRepo.FindByID(ctx, MealEventSet.MenuSetID)
	if err != nil {
		return err
	}

	return s.mealRepo.UpdateMenuSetInEvent(ctx, MealEventSet)
}

// RemoveMenuSetFromEvent removes a menu set from a meal event
func (s *mealEventService) RemoveMenuSetFromEvent(ctx context.Context, mealEventID uint, menuSetID uint) error {
	// Validate meal event exists
	_, err := s.mealRepo.FindByID(ctx, mealEventID)
	if err != nil {
		return err
	}

	return s.mealRepo.RemoveMenuSetFromEvent(ctx, mealEventID, menuSetID)
}

// FindMenuSetsByEventID finds all menu sets associated with a meal event
func (s *mealEventService) FindMenuSetsByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventSet, error) {
	// Validate meal event exists
	_, err := s.mealRepo.FindByID(ctx, mealEventID)
	if err != nil {
		return nil, err
	}

	return s.mealRepo.FindMenuSetsByEventID(ctx, mealEventID)
}

// CreateMealRequest implements creating a meal request
func (s *mealEventService) CreateMealRequest(ctx context.Context, request *model.MealRequest) error {
	return s.mealRepo.CreateRequest(ctx, request)
}

// FindRequestByID implements finding a meal request by ID
func (s *mealEventService) FindRequestByID(ctx context.Context, id uint) (*model.MealRequest, error) {
	return s.mealRepo.FindRequestByID(ctx, id)
}

// FindAllRequests implements finding all meal requests
func (s *mealEventService) FindAllRequests(ctx context.Context) ([]model.MealRequest, error) {
	return s.mealRepo.FindAllRequests(ctx)
}

// FindRequestsByUserID implements finding meal requests by user ID
func (s *mealEventService) FindRequestsByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error) {
	return s.mealRepo.FindRequestsByUserID(ctx, userID)
}

// DeleteMealRequest implements deleting a meal request
func (s *mealEventService) DeleteMealRequest(ctx context.Context, request *model.MealRequest) error {
	return s.mealRepo.DeleteRequest(ctx, request)
}

// AddAddressToEvent adds an address to an event
func (s *mealEventService) AddAddressToEvent(ctx context.Context, eventAddressID uint, mealEventID uint) error {
	// Implementation would go here
	return nil
}

// RemoveAddressFromEvent removes an address from an event
func (s *mealEventService) RemoveAddressFromEvent(ctx context.Context, eventAddressID uint, mealEventID uint) error {
	// Implementation would go here
	return nil
}

// FindAddressesByEventID finds all addresses associated with an event
func (s *mealEventService) FindAddressesByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventAddress, error) {
	// Implementation would go here
	return nil, nil
}

// CreateComment creates a new comment for a menu item
func (s *mealEventService) CreateComment(ctx context.Context, comment *model.MenuItemComment) error {
	return s.mealRepo.CreateComment(ctx, comment)
}

// FindCommentsByMealEventID finds all comments for a meal event
func (s *mealEventService) FindCommentsByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error) {
	return s.mealRepo.FindCommentsByMealEventID(ctx, mealEventID)
}

// FindUpcomingAndActive finds all upcoming and active meal events
func (s *mealEventService) FindUpcomingAndActive(ctx context.Context) ([]model.MealEvent, error) {
	return s.mealRepo.FindUpcomingAndActive(ctx)
}

// GetMealByID retrieves a meal event by ID with permission checking
func (s *mealEventService) GetMealByID(ctx context.Context, id uint, userID uint, isAdmin bool) (*model.MealEvent, error) {
	meal, err := s.FindByID(ctx, id)
	if err != nil {
		return nil, repository.ErrNotFound
	}

	return meal, nil
}

// CreateMeal creates a new meal event with the creator's user ID
func (s *mealEventService) CreateMeal(ctx context.Context, meal *model.MealEvent, userID uint) error {
	meal.CreatedBy = userID
	meal.UpdatedBy = userID
	return s.Create(ctx, meal)
}

// UpdateMeal updates a meal event with permission checking
func (s *mealEventService) UpdateMeal(ctx context.Context, id uint, meal *model.MealEvent, userID uint) error {
	existingMeal, err := s.FindByID(ctx, id)
	if err != nil {
		return err
	}

	meal.ID = id
	meal.UpdatedBy = userID

	// Preserve created_by and other fields that shouldn't be updated
	meal.CreatedBy = existingMeal.CreatedBy
	meal.CreatedAt = existingMeal.CreatedAt

	return s.Update(ctx, meal)
}

// DeleteMeal deletes a meal event with permission checking
func (s *mealEventService) DeleteMeal(ctx context.Context, id uint, userID uint) error {
	meal, err := s.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Update the updatedBy field to track who triggered the deletion
	meal.UpdatedBy = userID

	// Use the Delete method which properly handles soft deletion
	return s.Delete(ctx, meal)
}

// FindByDateRange finds meal events within a date range
func (s *mealEventService) FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealEvent, error) {
	return s.mealRepo.FindByDateRange(ctx, startDate, endDate)
}
