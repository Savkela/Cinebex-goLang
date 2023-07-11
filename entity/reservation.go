package entity

type Reservation struct {
	ProjectionId   int  `json:"projectionId"`
	UserId         int  `json:"userId"`
	TicketQuantity int  `json:"ticketQuantity"`
	TotalPrice     int  `json:"totalPrice"`
	Confirmed      bool `json:"confirmed"`
}
