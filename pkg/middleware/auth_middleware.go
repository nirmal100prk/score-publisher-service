package middleware

import (
	"net/http"
	"score-publisher-svc/pkg/jwtutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to validate JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		_, err := jwtutil.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// // Add the claims to the Gin context for use in handlers
		// c.Set("username", claims.Username)
		c.Next()
	}
}
