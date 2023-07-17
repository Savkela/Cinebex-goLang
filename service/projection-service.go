package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type ProjectionService interface {
	Save(projection entity.Projection) error
	FindAll() []entity.Projection
	FindOne(id string) (entity.Projection, error)
	Update(id string, Projection entity.Projection) (entity.Projection, error)
	Delete(id string) error
}

type projectionService struct {
	projections []entity.Projection
}

func NewProjectionService() ProjectionService {
	return &projectionService{}
}

func (service *projectionService) Save(projection entity.Projection) error {
	err := initializers.DB.Create(&projection).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *projectionService) FindOne(id string) (entity.Projection, error) {
	var projection entity.Projection
	result := initializers.DB.First(&projection, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return projection, errors.New("Projection not found")
		}
		return projection, result.Error
	}

	return projection, nil
}

func (service *projectionService) FindAll() []entity.Projection {
	var projections []entity.Projection
	initializers.DB.Find(&projections)
	return projections
}

func (service *projectionService) Update(id string, projection entity.Projection) (entity.Projection, error) {
	var projectionToUpdate entity.Projection
	err := initializers.DB.First(&projectionToUpdate, id).Error
	if err != nil {
		return projectionToUpdate, err
	}

	err = initializers.DB.Model(&projectionToUpdate).Updates(projection).Error
	if err != nil {
		return projectionToUpdate, err
	}

	return projectionToUpdate, nil
}

func (service *projectionService) Delete(id string) error {
	var projection entity.Projection
	result := initializers.DB.Delete(&projection, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
