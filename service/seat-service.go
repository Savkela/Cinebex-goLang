package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type SeatService interface {
	Save(seat entity.Seat) error
	FindOne(id string) (entity.Seat, error)
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

func (service *seatService) Save(seat entity.Seat) error {
	err := initializers.DB.Create(&seat).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *seatService) FindOne(id string) (entity.Seat, error) {
	var seat entity.Seat
	result := initializers.DB.First(&seat, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return seat, errors.New("Seat not found")
		}
		return seat, result.Error
	}

	return seat, nil
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
