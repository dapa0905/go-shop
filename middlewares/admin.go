package middlewares

import (
	"go-shop/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userRaw, exists := c.Get("user")
		log.Println("userRaw :", userRaw)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Requires authentication"})
			return
		}

		user := userRaw.(models.User)
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Administrator rights required"})
			return
		}

		c.Next()
	}
}
