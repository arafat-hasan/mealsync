package middleware

import (
	"errors"
	"fmt"

	"github.com/arafat-hasan/mealsync/internal/dto"
	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ErrorHandler middleware catches any panics and returns standardized error responses
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a request ID for tracking
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)

		// Process request
		c.Next()

		// If there was an error, it should have been set in c.Errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleError(c, err, requestID)
		}
	}
}

// handleError processes an error and returns an appropriate error response
func handleError(c *gin.Context, err error, requestID string) {
	var appErr *apperrors.AppError

	// Try to convert to an AppError
	if errors.As(err, &appErr) {
		// We have an AppError, add the request ID
		appErr.RequestID = requestID
	} else {
		// Create a new internal server error
		appErr = apperrors.NewInternalError("An unexpected error occurred", err).WithRequestID(requestID)
	}

	// Log error details (you could add more sophisticated logging here)
	fmt.Printf("Error [%s]: %s\n", requestID, appErr.Error())

	// Convert to DTO response
	response := dto.ErrorResponse{
		Error:   appErr.Message,
		Details: appErr.Details,
		Code:    string(appErr.Type),
	}

	// Return error response
	c.JSON(appErr.Code, response)
	c.Abort()
}

// HandleAppError is a helper for controllers to handle application errors
func HandleAppError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	requestID, _ := c.Get("RequestID")
	requestIDStr, _ := requestID.(string)

	handleError(c, err, requestIDStr)
}

// Recovery middleware handles panics
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID
				requestID, _ := c.Get("request_id")
				if requestID == nil {
					requestID = uuid.New().String()
				}

				var appErr *apperrors.AppError
				switch e := err.(type) {
				case error:
					appErr = apperrors.NewInternalError("Internal server error", e)
				default:
					appErr = apperrors.NewInternalError("Internal server error", errors.New("unknown panic"))
				}

				c.JSON(appErr.Code, appErr.WithRequestID(requestID.(string)))
				c.Abort()
			}
		}()
		c.Next()
	}
}
