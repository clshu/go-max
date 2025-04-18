package mw

import (
	"log"

	"github.com/gin-gonic/gin"
)

// LogUserAgent logs the details of each incoming request
func LogUserAgent(c *gin.Context) {
	ua := c.Request.UserAgent()
	log.Printf("User-Agent: %s | Method: %s | Path: %s\n", ua, c.Request.Method, c.Request.URL.Path)
	c.Next()
}
