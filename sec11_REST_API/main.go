package main

import (
	"net/http"
	"time"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.Use(mw.LogUserAgent)

	r.GET("/events", getEvents)
	r.POST("/events", createEvent)

	r.Run(":8080")

}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = len(models.GetAllEvents()) + 1 // Simple ID assignment
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	event.UserID = 1 // Assuming a default user ID for simplicity
	event.Save()
	c.JSON(http.StatusCreated, event)
}
