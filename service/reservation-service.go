package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type ReservationService interface {
	Save(entity.Reservation) entity.Reservation
	FindAll() []entity.Reservation
	FindOne(id string) entity.Reservation
}

type reservationService struct {
	reservations []entity.Reservation
}

func NewReservationService() ReservationService {
	return &reservationService{}
}

func (service *reservationService) Save(reservation entity.Reservation) entity.Reservation {
	service.reservations = append(service.reservations, reservation)
	return reservation
}

func (service *reservationService) FindOne(id string) entity.Reservation {
	var reservation entity.Reservation
	initializers.DB.First(&reservation, id)
	return reservation
}

func (service *reservationService) FindAll() []entity.Reservation {
	var reservations []entity.Reservation
	initializers.DB.Find(&reservations)
	return reservations
}
