package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type ReservationService interface {
	Save(reservation entity.Reservation) error
	FindAll() []entity.Reservation
	FindOne(id string) (entity.Reservation, error)
	Update(id string, reservation entity.Reservation) (entity.Reservation, error)
	Delete(id string) error
}

type reservationService struct {
	reservations []entity.Reservation
}

func NewReservationService() ReservationService {
	return &reservationService{}
}

func (service *reservationService) Save(reservation entity.Reservation) error {
	err := initializers.DB.Create(&reservation).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *reservationService) FindOne(id string) (entity.Reservation, error) {
	var reservation entity.Reservation
	result := initializers.DB.First(&reservation, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return reservation, errors.New("Reservation not found")
		}
		return reservation, result.Error
	}

	return reservation, nil
}

func (service *reservationService) FindAll() []entity.Reservation {
	var reservations []entity.Reservation
	initializers.DB.Find(&reservations)
	return reservations
}

func (service *reservationService) Update(id string, reservation entity.Reservation) (entity.Reservation, error) {
	var reservationToUpdate entity.Reservation
	err := initializers.DB.First(&reservationToUpdate, id).Error
	if err != nil {
		return reservationToUpdate, err
	}

	err = initializers.DB.Model(&reservationToUpdate).Updates(reservation).Error
	if err != nil {
		return reservationToUpdate, err
	}

	return reservationToUpdate, nil
}

func (service *reservationService) Delete(id string) error {
	var reservation entity.Reservation
	result := initializers.DB.Delete(&reservation, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
