package main

import (
	"log"

	"github.com/falasefemi2/chat-app/db"
	"github.com/falasefemi2/chat-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// gin router
	server := gin.Default()

	routes.RegisterRoute(server)

	// Start the server
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
