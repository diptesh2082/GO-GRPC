package main

import "fmt"

// import (
// 	"log"
// 	"github.com/gin-gonic/gin"
// 	"billing-software/cmd/internal/handlers/"
// )

// func main() {
// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define the route for billing
// 	router.GET("/billing", handlers.BillingHandler)

// 	// Start the server on port 8080
// 	log.Println("Starting server on :8080")
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal("Unable to start server: ", err)
// 	}
// }

type Logger interface {
    Log(message string)
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
    fmt.Println("Logging to file:", message)
}

func main() {
    var logger Logger = FileLogger{}
    logger.Log("Hello, Go!")
}
