package service

import (
	"errors"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService handles authentication-related operations
type AuthService struct {
	db *gorm.DB
}

// NewAuthService creates a new AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Register creates a new user
func (s *AuthService) Register(user *model.User) error {
	// Check if user already exists
	var existingUser model.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create user
	return s.db.Create(user).Error
}

// Authenticate verifies user credentials and returns the user if valid
func (s *AuthService) Authenticate(email, password string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// GenerateTokens creates a new pair of JWT tokens for a user
func (s *AuthService) GenerateTokens(user *model.User) (*TokenPair, error) {
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
		return nil, err
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
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// RefreshToken generates a new access token using a refresh token
func (s *AuthService) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	// Parse and validate refresh token
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_jwt_refresh_secret_key"), nil
	})

	if err != nil || !refreshToken.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Extract claims
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	// Get user ID from claims
	userID, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid user ID in refresh token")
	}

	// Get user from database
	var user model.User
	if err := s.db.First(&user, uint(userID)).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new tokens
	return s.GenerateTokens(&user)
}

// TokenPair represents a pair of JWT tokens
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
