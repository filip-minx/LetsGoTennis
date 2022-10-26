package main

import (
	"fmt"
	"hackathon/api/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// represents data about a courts and reservation.
type courts struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type reservation struct {
	Phone       string  `json:"id"`
	PeopleCount string  `json:"name"`
	StartTime   float64 `json:"price"`
	EndTime     float64 `json:"price"`
}

var courtsList = []courts{
	{ID: "1", Name: "Grass", Price: 9.99},
	{ID: "2", Name: "Concrete", Price: 9.99},
	{ID: "3", Name: "Clay", Price: 9.99},
	{ID: "1", Name: "Ice", Price: 19.99},
}

// getCourts responds with the list of all courts as JSON.
func getCourts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courtsList)
}

func main() {

	router := gin.Default()
	router.GET("/courts", getCourts)

	router.Run("localhost:8080")

	fmt.Println("Hello, World!")
	fmt.Println(api.Hello("asd"))

}
