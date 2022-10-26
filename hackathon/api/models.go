package api

// represents data about a courts and reservation.
type Courts struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Reservation struct {
	Phone       float64 `json:"phone"`
	PlayerCount float64 `json:"playerCount"`
	StartTime   float64 `json:"startTime"`
	EndTime     float64 `json:"endTime"`
	CourtName   string  `json:"courtName"`
}

var CourtsList = []Courts{
	{ID: "1", Name: "Grass", Price: 9.99},
	{ID: "2", Name: "Concrete", Price: 9.99},
	{ID: "3", Name: "Clay", Price: 9.99},
	{ID: "1", Name: "Ice", Price: 19.99},
}

var Reservations []Reservation
