package api

import (
	"net/http"

	"github.com/arafat-hasan/mealsync/internal/dto"
	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/mapper"
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register handles user registration
//	@Summary		Register new user
//	@Description	Create a new user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserRegisterRequest	true	"User registration data"
//	@Success		201		{object}	dto.UserResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		409		{object}	dto.ErrorResponse
//	@Router			/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	user := &model.User{
		EmployeeID: req.EmployeeID,
		Username:   req.Username,
		Password:   req.Password,
		Name:       req.Name,
		Email:      req.Email,
		Department: req.Department,
		Role:       model.UserRoleEmployee, // Default role
	}

	if err := h.authService.Register(user); err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	c.JSON(http.StatusCreated, mapper.ToUserResponse(user))
}

// Login handles user login
//	@Summary		Login user
//	@Description	Authenticate user and return JWT tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserLoginRequest	true	"Login credentials"
//	@Success		200		{object}	dto.TokenResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	// Authenticate user
	user, err := h.authService.Authenticate(req.Email, req.Password)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Invalid credentials", err))
		return
	}

	// Generate tokens
	tokens, err := h.authService.GenerateTokens(user)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewInternalError("Failed to generate tokens", err))
		return
	}

	// Set custom header for Swagger UI
	c.Header("X-Access-Token", tokens.AccessToken)

	// Return tokens in response body
	response := mapper.ToTokenResponse(tokens.AccessToken, tokens.RefreshToken, user)
	c.JSON(http.StatusOK, response)
}

// RefreshToken handles token refresh
//	@Summary		Refresh access token
//	@Description	Get a new access token using a refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RefreshTokenRequest	true	"Refresh token"
//	@Success		200		{object}	service.TokenPair
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.HandleAppError(c, apperrors.NewValidationError("Invalid request format", err))
		return
	}

	tokens, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		middleware.HandleAppError(c, apperrors.NewUnauthorizedError("Invalid refresh token", err))
		return
	}

	// Set custom header for Swagger UI
	c.Header("X-Access-Token", tokens.AccessToken)

	// Return tokens in response body
	c.JSON(http.StatusOK, tokens)
}
