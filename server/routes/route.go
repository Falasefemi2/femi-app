package routes

import (
	"github.com/falasefemi2/chat-app/conrollers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.POST("/users", conrollers.CreateUser)
	server.DELETE("/users", conrollers.DeleteAllUsers) // Route to delete all users
	server.POST("/login", conrollers.Login)
	server.POST("/channels", conrollers.CreateChannel)
}
