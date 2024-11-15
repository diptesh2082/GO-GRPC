package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"billing-software/cmd/internal/handlers"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define the route for billing
	router.GET("/billing", handlers.BillingHandler)

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
