package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicBindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}