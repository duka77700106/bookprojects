package middleware

import (
	"bookProject/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			return
		}

		c.Set("userID", uint(userID))

		role, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid role"})
			return
		}
		c.Set("role", role)

		c.Next()
	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not set"})
			return
		}
		role, ok := raw.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role type"})
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}
