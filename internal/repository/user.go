package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// userRepository implements UserRepository interface
type userRepository struct {
	*baseRepository[model.User]
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		baseRepository: NewBaseRepository[model.User](db),
	}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.baseRepository.Create(ctx, user)
}

// FindByID finds a user by ID
func (r *userRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all users
func (r *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	return r.baseRepository.FindAll(ctx)
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.baseRepository.Update(ctx, user)
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.baseRepository.Delete(ctx, id)
}

// FindByEmail finds a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.baseRepository.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername finds a user by username
func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.baseRepository.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
