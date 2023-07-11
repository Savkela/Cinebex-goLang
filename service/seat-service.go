package service

import "cinebex/entity"

type SeatService interface {
	Save(entity.Seat) entity.Seat
	FindAll() []entity.Seat
}

type seatService struct {
	seats []entity.Seat
}

func NewSeatService() SeatService {
	return &seatService{}
}

func (service *seatService) Save(seat entity.Seat) entity.Seat {
	service.seats = append(service.seats, seat)
	return seat
}

func (service *seatService) FindAll() []entity.Seat {
	return service.seats
}
