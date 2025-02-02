package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	s.GET("/events", getEvents)

	if err := s.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello"})
}
