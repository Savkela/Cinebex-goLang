package entity

type Projection struct {
	MovieId        int    `json:"movieId"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	TicketPrice    string `json:"ticketPrice"`
	AvailableSeats string `json:"availableSeats"`
}
