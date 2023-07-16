package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type SeatService interface {
	Save(entity.Seat) entity.Seat
	FindOne(id string) entity.Seat
	FindAll() []entity.Seat
	Update(id string, seat entity.Seat) (entity.Seat, error)
	Delete(id string) error
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

func (service *seatService) FindOne(id string) entity.Seat {
	var seat entity.Seat
	initializers.DB.First(&seat, id)
	return seat
}

func (service *seatService) FindAll() []entity.Seat {
	var seats []entity.Seat
	initializers.DB.Find(&seats)
	return seats
}

func (service *seatService) Update(id string, seat entity.Seat) (entity.Seat, error) {
	var seatToUpdate entity.Seat
	err := initializers.DB.First(&seatToUpdate, id).Error
	if err != nil {
		return seatToUpdate, err
	}

	err = initializers.DB.Model(&seatToUpdate).Updates(seat).Error
	if err != nil {
		return seatToUpdate, err
	}

	return seatToUpdate, nil
}

func (service *seatService) Delete(id string) error {
	var seat entity.Seat
	result := initializers.DB.Delete(&seat, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
