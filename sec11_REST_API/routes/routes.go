package routes

import (
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)

	protected := r.Group("/")
	protected.Use(middleware.Authenticate)

	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)

	r.POST("/signup", createUser)
	r.POST("/login", loginUser)

}
