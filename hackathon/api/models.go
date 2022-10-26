package api

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

var CourtsList = []courts{
	{ID: "1", Name: "Grass", Price: 9.99},
	{ID: "2", Name: "Concrete", Price: 9.99},
	{ID: "3", Name: "Clay", Price: 9.99},
	{ID: "1", Name: "Ice", Price: 19.99},
}
