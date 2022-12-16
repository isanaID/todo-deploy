package middlewares

import (
	"context"
	"database/sql"
	"strings"
	"todo/controllers"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing auth token"})
			return
		}

		auth := strings.Split(authHeader, " ")
		userData, err := controllers.DecryptJWT(auth[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid auth token"})
			return
		}

		ctxUserID := context.WithValue(c.Request.Context(), "userID", userData["UserID"])
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}