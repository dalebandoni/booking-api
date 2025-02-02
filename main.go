package main

import (
	"log"
	"net/http"

	"github.com/dalebandoni/booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	s.GET("/events", getEvents)
	s.POST("/events", createEvent)

	if err := s.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEvents(context *gin.Context) {
	e := models.GetAllEvents()
	context.JSON(http.StatusOK, e)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}
