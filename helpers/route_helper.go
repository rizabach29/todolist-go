package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(&obj); err != nil {
		return err
	}
	return nil
}

func IsRoleAllowed(c *gin.Context, allowedRole string) bool {
	role, _ := c.Get("role")
	if role != allowedRole {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "role must be" + allowedRole})
		return false
	}
	return true
}