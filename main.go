package main

import (
	"log"
	"net/http"

	"github.com/dalebandoni/booking-api/db"
	"github.com/dalebandoni/booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	s := gin.Default()

	s.GET("/events", getEvents)
	s.POST("/events", createEvent)

	if err := s.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEvents(context *gin.Context) {
	e, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later."})
		return
	}
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}
