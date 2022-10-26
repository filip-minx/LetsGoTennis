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

	// Call BindJSON to bid the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newReservation); err != nil {
		return
	}

	var durationInMs = newReservation.EndTime - newReservation.StartTime
	var durationInMins = float64(durationInMs) / float64((60 * 60))
	court := FirstOrDefault(CourtsList, func(p *Court) bool {
		return p.Name == newReservation.CourtName
	})
	var price = court.Price * float64(durationInMins)

	if newReservation.PlayerCount > 3 {
		price = price * 1.5
	}

	newReservation.Price = price

	// Add the new album to the slice.
	Reservations = append(Reservations, newReservation)
	c.IndentedJSON(http.StatusCreated, newReservation)

}

func GetReservations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Reservations)
}

func FirstOrDefault[T any](slice []T, filter func(*T) bool) (element *T) {

	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			return &slice[i]
		}
	}

	return nil
}
