package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)
	r.POST("/events", createEvent)
}
