package routes

import (
	"net/http"
	"strconv"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userID := c.GetInt64("userID")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	foundEvent, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find event"})
		return
	}

	err = foundEvent.Register(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
	})

}

func cancelRegistration(c *gin.Context) {
	userID := c.GetInt64("userID")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventID

	// Check if the registration exists

	found, err := event.CheckRegistration(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check registration"})
		return
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No registration found"})
		return
	}

	err = event.CancelRegistration(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled",
	})

}
