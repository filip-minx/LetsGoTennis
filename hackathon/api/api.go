package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CourtsList)
}

func GetReservations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Reservations)
}
