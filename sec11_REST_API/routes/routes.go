package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)
	r.POST("/events", createEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
	r.POST("/signup", createUser)

}
