package main

import (
	"log"

	"github.com/dalebandoni/booking-api/db"
	"github.com/dalebandoni/booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	s := gin.Default()

	routes.RegisterRoutes(s)

	if err := s.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
