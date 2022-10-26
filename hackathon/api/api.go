package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCourts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CourtsList)
}

func PostReservation(c *gin.Context) {
	var newReservation Reservation

	if err := c.BindJSON(&newReservation); err != nil {
		return
	}
	newReservation.ID = newReservation.StartTime
	court := FirstOrDefault(CourtsList, func(p *Court) bool {
		return p.Name == newReservation.CourtName
	})

	if !IsCourtAvailable(court, newReservation.StartTime, newReservation.EndTime) {
		c.IndentedJSON(http.StatusTooEarly, newReservation)
	}

	var durationInMs = newReservation.EndTime - newReservation.StartTime
	var durationInMins = float64(durationInMs) / float64((60 * 60))

	var price = court.Price * float64(durationInMins)

	if newReservation.PlayerCount > 3 {
		price = price * 1.5
	}

	newReservation.Price = price

	// Add the new album to the slice.
	Reservations = append(Reservations, newReservation)
	c.IndentedJSON(http.StatusCreated, newReservation)
}

func IsCourtAvailable(court *Court, startTime int64, endTime int64) bool {
	var reservations = Where(Reservations, func(r *Reservation) bool {
		return r.CourtName == court.Name
	})

	var res = FirstOrDefault(reservations, func(r *Reservation) bool {
		return endTime > r.StartTime && startTime < r.EndTime
	})

	return res == nil
}

func GetReservations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Reservations)
}

func GetReservationId(c *gin.Context) {
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

func Where[T any](slice []T, filter func(*T) bool) []T {

	var ret []T = make([]T, 0)

	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			ret = append(ret, slice[i])
		}
	}

	return ret
}

func DeleteReservation(c *gin.Context) {
	id32, err := strconv.Atoi(c.Param("id"))
	var id = int64(id32)

	fmt.Println(err)

	var matchingReservation = FirstOrDefault(Reservations, func(r *Reservation) bool {
		return id == r.ID
	})

	var deleteIndex = IndexOf(Reservations, *matchingReservation)

	Reservations = remove(Reservations, deleteIndex)

}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}
