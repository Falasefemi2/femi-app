package conrollers

import (
	"net/http"

	"github.com/falasefemi2/chat-app/models"
	"github.com/gin-gonic/gin"
)

func CreateChannel(c *gin.Context) {
	var reqBody struct {
		Name      string `json:"name"`
		CreatedBy int    `json:"createdBy"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel, err := models.CreateChannel(reqBody.Name, reqBody.CreatedBy)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, channel)

}
