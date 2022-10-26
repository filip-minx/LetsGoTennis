package main

import (
	"hackathon/api/api"

	"github.com/gin-gonic/gin"
)

// getCourts responds with the list of all courts as JSON.

func main() {

	router := gin.Default()
	router.POST("api/reservations", api.PostReservation)
	router.GET("api/reservations", api.GetReservations)
	router.GET("api/courts", api.GetCourts)

	router.Run("localhost:8080")
}
