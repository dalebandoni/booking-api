package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(s *gin.Engine) {
	s.GET("/events", getEvents)
	s.GET("/events/:id", getEvent)
	s.POST("/events", createEvent)
	s.PUT("/events/:id", updateEvent)
	s.DELETE("/events/:id", deleteEvent)

	s.POST("/signup", signup)
}
