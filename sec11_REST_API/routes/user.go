package routes

import (
	"net/http"
	"strings"

	"example.com/restapi/models"
	"example.com/restapi/util"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := user.Save()
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed") {
			errMsg = "email existed"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func loginUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user credentials, it also populates u.ID
	if err := user.ValidateCredentials(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
