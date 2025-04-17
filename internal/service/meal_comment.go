package service

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// mealCommentService handles business logic for meal comment operations
type mealCommentService struct {
	commentRepo repository.MealCommentRepository
	mealRepo    repository.MealEventRepository
	userRepo    repository.UserRepository
}

// NewMealCommentService creates a new instance of MealCommentService
func NewMealCommentService(
	commentRepo repository.MealCommentRepository,
	mealRepo repository.MealEventRepository,
	userRepo repository.UserRepository,
) MealCommentService {
	return &mealCommentService{
		commentRepo: commentRepo,
		mealRepo:    mealRepo,
		userRepo:    userRepo,
	}
}

// GetComments retrieves comments for a meal event
func (s *mealCommentService) GetComments(ctx context.Context, mealEventID uint) ([]model.MealComment, error) {
	// Verify meal event exists
	_, err := s.mealRepo.FindByID(ctx, mealEventID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.FindByMealEventID(ctx, mealEventID)
}

// GetCommentByID retrieves a specific comment by ID
func (s *mealCommentService) GetCommentByID(ctx context.Context, id uint) (*model.MealComment, error) {
	return s.commentRepo.FindByID(ctx, id)
}

// CreateComment creates a new comment
func (s *mealCommentService) CreateComment(ctx context.Context, comment *model.MealComment, userID uint) error {
	if comment == nil {
		return errors.NewValidationError("comment cannot be nil", nil)
	}

	// Verify meal event exists
	_, err := s.mealRepo.FindByID(ctx, comment.MealEventID)
	if err != nil {
		return err
	}

	// Set comment fields
	comment.UserID = userID
	comment.CreatedBy = userID
	comment.UpdatedBy = userID

	return s.commentRepo.Create(ctx, comment)
}

// UpdateComment updates an existing comment
func (s *mealCommentService) UpdateComment(ctx context.Context, id uint, comment *model.MealComment, userID uint) error {
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
	existingComment.UpdatedBy = userID

	return s.commentRepo.Update(ctx, existingComment)
}

// DeleteComment soft deletes a comment
func (s *mealCommentService) DeleteComment(ctx context.Context, id uint, userID uint) error {
	comment, err := s.commentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Verify ownership
	if comment.UserID != userID {
		return errors.NewForbiddenError("unauthorized to delete this comment", nil)
	}

	comment.UpdatedBy = userID
	return s.commentRepo.Delete(ctx, comment)
}

// GetReplies retrieves replies for a comment
func (s *mealCommentService) GetReplies(ctx context.Context, commentID uint) ([]model.MealComment, error) {
	// Verify parent comment exists
	_, err := s.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.FindByParentCommentID(ctx, commentID)
}

// GetUserComments retrieves all comments by a user
func (s *mealCommentService) GetUserComments(ctx context.Context, userID uint) ([]model.MealComment, error) {
	return s.commentRepo.FindByUserID(ctx, userID)
}
