package api

// represents data about a courts and reservation.
type courts struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Reservation struct {
	Phone       float64 `json:"id"`
	PeopleCount float64 `json:"name"`
	StartTime   float64 `json:"price"`
	EndTime     float64 `json:"price"`
}

var CourtsList = []courts{
	{ID: "1", Name: "Grass", Price: 9.99},
	{ID: "2", Name: "Concrete", Price: 9.99},
	{ID: "3", Name: "Clay", Price: 9.99},
	{ID: "1", Name: "Ice", Price: 19.99},
}

var Reservations []Reservation
