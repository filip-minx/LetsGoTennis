package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CourtsList)
}

func PostReservation(c *gin.Context) {
	var newReservation Reservation

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newReservation); err != nil {
		return
	}

	// Add the new album to the slice.
	Reservations = append(Reservations, newReservation)
	c.IndentedJSON(http.StatusCreated, newReservation)
}
