package entity

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	ReservationId int64 `json:"reservationId"`
	Rating        int64 `json:"rating"`
}
