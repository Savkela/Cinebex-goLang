package entity

type Reservation struct {
	ID             int64    `json:"id"`
	ProjectionId   int64    `json:"projectionId"`
	UserId         int64    `json:"userId"`
	TicketQuantity int64    `json:"ticketQuantity"`
	TotalPrice     float64  `json:"totalPrice"`
	Confirmed      bool     `json:"confirmed"`
	Seats          []Seat   `json:"seats"`
	Ratings        []Rating `json:"ratings"`
}
