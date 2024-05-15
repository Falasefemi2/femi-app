package main

import (
	"log"

	"github.com/falasefemi2/chat-app/conrollers"
	"github.com/falasefemi2/chat-app/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// gin router
	router := gin.Default()

	// set up routes
	router.POST("/users", conrollers.CreateUser)
	router.DELETE("/users", conrollers.DeleteAllUsers) // Route to delete all users

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
