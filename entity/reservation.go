package entity

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ProjectionId   int64    `json:"projectionId"`
	UserId         int64    `json:"userId"`
	TicketQuantity int64    `json:"ticketQuantity"`
	TotalPrice     float64  `json:"totalPrice"`
	Confirmed      bool     `json:"confirmed"`
	Seats          []Seat   `json:"seats"`
	Ratings        []Rating `json:"ratings"`
}
