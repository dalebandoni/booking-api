package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	s.GET("/events", getEvents)

	s.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello"})
}
