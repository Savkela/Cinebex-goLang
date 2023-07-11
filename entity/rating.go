package entity

type Rating struct {
	ID            int64 `json:"id"`
	ReservationId int64 `json:"reservationId"`
	Rating        int64 `json:"rating"`
}
