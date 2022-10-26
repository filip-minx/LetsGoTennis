package main

import (
	"hackathon/api/api"

	"github.com/gin-gonic/gin"
)

// getCourts responds with the list of all courts as JSON.

func main() {

	router := gin.Default()
	router.GET("api/courts", api.GetCourts)
	router.GET("api/reservations", api.GetReservations)

	api.Reservations = append(
		api.Reservations,
		api.Reservation{Phone: 1.0, PeopleCount: 1.0, StartTime: 1.0, EndTime: 1.0})

	router.Run("localhost:8080")

}
