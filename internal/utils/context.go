package utils

import (
	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext extracts user ID from the context set by auth middleware
func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.NewUnauthorizedError("user not authenticated", nil)
	}
	return userID.(uint), nil
}

// IsAdminFromContext checks if the user is an admin based on role in context
func IsAdminFromContext(c *gin.Context) bool {
	role, exists := c.Get("role")
	if !exists {
		return false // This should not happen if auth middleware is working correctly
	}
	return role.(string) == "admin"
}
