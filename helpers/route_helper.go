package helpers

import (
	"github.com/gin-gonic/gin"
)

func PanicBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(&obj); err != nil {
		return err
	}
	return nil
}