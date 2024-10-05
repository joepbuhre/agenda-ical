package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenInput string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != tokenInput {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
