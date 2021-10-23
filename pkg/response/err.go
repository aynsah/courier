package response

import (
	"github.com/gin-gonic/gin"
)

func ShowErr(c *gin.Context, status int, message string, data interface{}) {
	// Menyimpan error kedalam log bila data bukan nil
	if data != nil {
		LogErr(c.FullPath(), status, message, data)
	}
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
	})
	c.Abort()
}
