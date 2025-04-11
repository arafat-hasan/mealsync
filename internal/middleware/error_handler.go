package middleware

import (
	"errors"
	"net/http"

	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ErrorHandler middleware handles errors consistently
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		// Process request
		c.Next()

		// Check for errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Handle different types of errors
			var appErr *apperrors.AppError
			if errors.As(err, &appErr) {
				// AppError is already formatted
				c.JSON(appErr.Code, appErr.WithRequestID(requestID))
				return
			}

			// Handle GORM errors
			if errors.Is(err, gorm.ErrRecordNotFound) {
				appErr = apperrors.NewNotFoundError("Resource not found", err)
				c.JSON(appErr.Code, appErr.WithRequestID(requestID))
				return
			}

			// Handle other database errors
			if err == gorm.ErrInvalidTransaction ||
				err == gorm.ErrNotImplemented ||
				err == gorm.ErrMissingWhereClause ||
				err == gorm.ErrUnsupportedRelation {
				appErr = apperrors.NewInternalError("Database error", err)
				c.JSON(appErr.Code, appErr.WithRequestID(requestID))
				return
			}

			// Handle unknown errors
			appErr = apperrors.NewInternalError("Internal server error", err)
			c.JSON(appErr.Code, appErr.WithRequestID(requestID))
		}

		// If no errors, continue with the response
		if c.Writer.Status() == http.StatusOK {
			c.Next()
		}
	}
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
