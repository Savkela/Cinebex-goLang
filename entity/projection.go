package entity

import (
	"time"

	"gorm.io/gorm"
)

type Projection struct {
	gorm.Model
	MovieId             int64         `json:"movie" binding:"required"`
	Date                time.Time     `json:"date"`
	Time                time.Time     `json:"time"`
	Price               float64       `json:"ticketPrice"`
	TotalAvailableSeats int64         `json:"availableSeats"`
	Reservations        []Reservation `json:"reservations"`
}
