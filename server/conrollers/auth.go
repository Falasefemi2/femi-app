package conrollers

import (
	"net/http"

	"github.com/falasefemi2/chat-app/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.AuthenticateUser(credentials.Email, credentials.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Remove password field from the user object for security
	user.Password = ""

	// c.JSON(http.StatusOK, gin.H{"user": user})
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}
