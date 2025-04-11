package service

import (
	"errors"
	"time"

	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService handles authentication-related operations
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(db),
	}
}

// Register creates a new user
func (s *AuthService) Register(user *model.User) error {
	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(nil, user.Email)
	if err == nil && existingUser != nil {
		return apperrors.NewConflictError("User already exists", nil)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.NewInternalError("Failed to hash password", err)
	}
	user.Password = string(hashedPassword)

	// Create user
	if err := s.userRepo.Create(nil, user); err != nil {
		return apperrors.NewInternalError("Failed to create user", err)
	}

	return nil
}

// Authenticate verifies user credentials and returns the user if valid
func (s *AuthService) Authenticate(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(nil, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewUnauthorizedError("Invalid credentials", nil)
		}
		return nil, apperrors.NewInternalError("Failed to find user", err)
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, apperrors.NewUnauthorizedError("Invalid credentials", nil)
	}

	return user, nil
}

// GenerateTokens creates a new pair of JWT tokens for a user
func (s *AuthService) GenerateTokens(user *model.User) (*TokenPair, error) {
	if user == nil {
		return nil, apperrors.NewValidationError("User is required", nil)
	}

	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
		"role": user.Role,
	})

	// Sign access token
	accessTokenString, err := accessToken.SignedString([]byte("your_jwt_secret_key"))
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to generate access token", err)
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	})

	// Sign refresh token
	refreshTokenString, err := refreshToken.SignedString([]byte("your_jwt_refresh_secret_key"))
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to generate refresh token", err)
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// RefreshToken generates a new access token using a refresh token
func (s *AuthService) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	if refreshTokenString == "" {
		return nil, apperrors.NewValidationError("Refresh token is required", nil)
	}

	// Parse and validate refresh token
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_jwt_refresh_secret_key"), nil
	})

	if err != nil || !refreshToken.Valid {
		return nil, apperrors.NewUnauthorizedError("Invalid refresh token", err)
	}

	// Extract claims
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apperrors.NewUnauthorizedError("Invalid refresh token claims", nil)
	}

	// Get user ID from claims
	userID, ok := claims["sub"].(float64)
	if !ok {
		return nil, apperrors.NewUnauthorizedError("Invalid user ID in refresh token", nil)
	}

	// Get user from database
	user, err := s.userRepo.FindByID(nil, uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewUnauthorizedError("User not found", nil)
		}
		return nil, apperrors.NewInternalError("Failed to find user", err)
	}

	// Generate new tokens
	return s.GenerateTokens(user)
}

// TokenPair represents a pair of JWT tokens
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
