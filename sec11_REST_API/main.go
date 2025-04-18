package main

import (
	"net/http"

	"example.com/restapi/db"
	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	// r.Use(mw.LogUserAgent)

	r.GET("/events", getEvents)
	r.POST("/events", createEvent)

	r.Run(":8080")

}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = 1 // Assuming a default user ID for simplicity

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}
