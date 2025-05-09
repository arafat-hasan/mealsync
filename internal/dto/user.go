package dto

import "github.com/arafat-hasan/mealsync/internal/model"

// UserResponse represents user data in responses
type UserResponse struct {
	BaseResponse
	EmployeeID          string         `json:"employee_id"`
	Username            string         `json:"username"`
	Name                string         `json:"name"`
	Email               string         `json:"email"`
	Department          string         `json:"department"`
	Role                model.UserRole `json:"role"`
	NotificationEnabled bool           `json:"notification_enabled"`
	IsActive            bool           `json:"is_active"`
}

// UserCreateRequest represents data for creating a new user
type UserCreateRequest struct {
	EmployeeID          string         `json:"employee_id" binding:"required"`
	Username            string         `json:"username" binding:"required"`
	Password            string         `json:"password" binding:"required,min=6"`
	Name                string         `json:"name" binding:"required"`
	Email               string         `json:"email" binding:"required,email"`
	Department          string         `json:"department" binding:"required"`
	Role                model.UserRole `json:"role"`
	NotificationEnabled bool           `json:"notification_enabled"`
}

// UserUpdateRequest represents data for updating an existing user
type UserUpdateRequest struct {
	Username            *string         `json:"username"`
	Password            *string         `json:"password" binding:"omitempty,min=6"`
	Name                *string         `json:"name"`
	Email               *string         `json:"email" binding:"omitempty,email"`
	Department          *string         `json:"department"`
	Role                *model.UserRole `json:"role"`
	NotificationEnabled *bool           `json:"notification_enabled"`
	IsActive            *bool           `json:"is_active"`
}

// UserLoginRequest represents user login data
type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserRegisterRequest represents user registration data
type UserRegisterRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Department string `json:"department" binding:"required"`
}

// TokenResponse represents authentication token response
type TokenResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

// RefreshTokenRequest represents a token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
