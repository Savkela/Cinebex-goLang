package service

import "cinebex/entity"

type ReservationService interface {
	Save(entity.Reservation) entity.Reservation
	FindAll() []entity.Reservation
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

func (service *reservationService) FindAll() []entity.Reservation {
	return service.reservations
}
