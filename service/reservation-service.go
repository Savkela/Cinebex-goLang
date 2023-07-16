package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type ReservationService interface {
	Save(entity.Reservation) entity.Reservation
	FindAll() []entity.Reservation
	FindOne(id string) entity.Reservation
	Update(id string, reservation entity.Reservation) (entity.Reservation, error)
	Delete(id string) error
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
