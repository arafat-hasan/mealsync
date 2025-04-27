package service

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// menuItemCommentService handles business logic for menu item comment operations
type menuItemCommentService struct {
	commentRepo repository.MenuItemCommentRepository
	mealRepo    repository.MealEventRepository
	userRepo    repository.UserRepository
	menuRepo    repository.MenuItemRepository
}

// NewMenuItemCommentService creates a new instance of MenuItemCommentService
func NewMenuItemCommentService(
	commentRepo repository.MenuItemCommentRepository,
	mealRepo repository.MealEventRepository,
	userRepo repository.UserRepository,
	menuRepo repository.MenuItemRepository,
) MenuItemCommentService {
	return &menuItemCommentService{
		commentRepo: commentRepo,
		mealRepo:    mealRepo,
		userRepo:    userRepo,
		menuRepo:    menuRepo,
	}
}

// GetComments retrieves comments for a meal event
func (s *menuItemCommentService) GetComments(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error) {
	// Verify meal event exists
	_, err := s.mealRepo.FindByID(ctx, mealEventID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.FindByMealEventID(ctx, mealEventID)
}

// GetCommentByID retrieves a specific comment by ID
func (s *menuItemCommentService) GetCommentByID(ctx context.Context, id uint) (*model.MenuItemComment, error) {
	return s.commentRepo.FindByID(ctx, id)
}

// CreateComment creates a new comment
func (s *menuItemCommentService) CreateComment(ctx context.Context, comment *model.MenuItemComment, userID uint) error {
	if comment == nil {
		return errors.NewValidationError("comment cannot be nil", nil)
	}

	// Verify meal event exists
	_, err := s.mealRepo.FindByID(ctx, comment.MealEventID)
	if err != nil {
		return err
	}

	// Verify menu item exists
	_, err = s.menuRepo.FindByID(ctx, comment.MenuItemID)
	if err != nil {
		return err
	}

	// Set comment fields
	comment.UserID = userID

	comment.CreatedByID = userID
	comment.UpdatedByID = userID

	return s.commentRepo.Create(ctx, comment)
}

// UpdateComment updates an existing comment
func (s *menuItemCommentService) UpdateComment(ctx context.Context, id uint, comment *model.MenuItemComment, userID uint) error {
	if comment == nil {
		return errors.NewValidationError("comment cannot be nil", nil)
	}

	existingComment, err := s.commentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership
	if existingComment.UserID != userID {
		return errors.NewForbiddenError("unauthorized to update this comment", nil)
	}

	// Update fields
	existingComment.Comment = comment.Comment
	existingComment.Rating = comment.Rating

	existingComment.UpdatedByID = userID

	return s.commentRepo.Update(ctx, existingComment)
}

// DeleteComment soft deletes a comment
func (s *menuItemCommentService) DeleteComment(ctx context.Context, id uint, userID uint) error {
	comment, err := s.commentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership
	if comment.UserID != userID {
		return errors.NewForbiddenError("unauthorized to delete this comment", nil)
	}

	comment.UpdatedByID = userID
	return s.commentRepo.Delete(ctx, comment)
}

// GetUserComments retrieves all comments by a user
func (s *menuItemCommentService) GetUserComments(ctx context.Context, userID uint) ([]model.MenuItemComment, error) {
	return s.commentRepo.FindByUserID(ctx, userID)
}

// GetMenuItemComments retrieves all comments for a specific menu item
func (s *menuItemCommentService) GetMenuItemComments(ctx context.Context, menuItemID uint) ([]model.MenuItemComment, error) {
	// Verify menu item exists
	_, err := s.menuRepo.FindByID(ctx, menuItemID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.FindByMenuItemID(ctx, menuItemID)
}

// GetReplies retrieves all replies to a specific comment
func (s *menuItemCommentService) GetReplies(ctx context.Context, commentID uint) ([]model.MenuItemComment, error) {
	// Verify parent comment exists
	_, err := s.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	// Fetch all replies to this comment
	return s.commentRepo.FindReplies(ctx, commentID)
}
