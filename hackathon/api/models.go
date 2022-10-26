package api

// represents data about a courts and reservation.
type Court struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Reservation struct {
	ID          int64   `json:"id"`
	Phone       float64 `json:"phone"`
	PlayerCount float64 `json:"playerCount"`
	StartTime   int64   `json:"startTime"`
	EndTime     int64   `json:"endTime"`
	CourtName   string  `json:"courtName"`
	Price       float64 `json:"price"`
}

var CourtsList = []Court{
	{ID: "1", Name: "Grass", Price: 9.99},
	{ID: "2", Name: "Concrete", Price: 9.99},
	{ID: "3", Name: "Clay", Price: 9.99},
	{ID: "4", Name: "Ice", Price: 19.99},
}

var Reservations []Reservation
