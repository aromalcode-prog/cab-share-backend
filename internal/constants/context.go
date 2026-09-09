package constants

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	UserIDKey = "userID"
)

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, errors.New("user ID not found in context")
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		return 0, errors.New("user ID is not a uint")
	}
	return userIDUint, nil
}
