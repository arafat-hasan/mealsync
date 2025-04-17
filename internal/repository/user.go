package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// userRepository implements UserRepository interface
type userRepository struct {
	*baseRepository[model.User]
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		baseRepository: NewBaseRepository[model.User](db),
		db:             db,
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

// FindActive finds all active users
func (r *userRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.User, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.baseRepository.Update(ctx, user)
}

// Delete deletes a user (soft delete)
func (r *userRepository) Delete(ctx context.Context, user *model.User) error {
	return r.baseRepository.Delete(ctx, user)
}

// HardDelete permanently deletes a user
func (r *userRepository) HardDelete(ctx context.Context, user *model.User) error {
	return r.baseRepository.HardDelete(ctx, user)
}

// FindByEmail finds a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername finds a user by username
func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmployeeID finds a user by employee ID
func (r *userRepository) FindByEmployeeID(ctx context.Context, employeeID int) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("employee_id = ?", employeeID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
