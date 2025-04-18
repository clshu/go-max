package middleware

import (
	"net/http"
	"strings"

	"example.com/restapi/util"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware function that checks for a valid JWT token in the request header
func Authenticate(c *gin.Context) {
	// Get the token from the request header
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	if !strings.HasPrefix(token, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}
	// Remove the "Bearer " prefix
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	userID, err := util.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Store the user ID in the context for later use
	c.Set("userID", userID)
	// Proceed to the next middleware or handler
	c.Next()
}
