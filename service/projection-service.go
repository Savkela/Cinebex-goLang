package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type ProjectionService interface {
	Save(entity.Projection) entity.Projection
	FindAll() []entity.Projection
	FindOne(id string) entity.Projection
	Update(id string, user entity.Projection) (entity.Projection, error)
	Delete(id string) error
}

type projectionService struct {
	projections []entity.Projection
}

func NewProjectionService() ProjectionService {
	return &projectionService{}
}

func (service *projectionService) Save(projection entity.Projection) entity.Projection {
	service.projections = append(service.projections, projection)
	return projection
}

func (service *projectionService) FindOne(id string) entity.Projection {
	var projection entity.Projection
	initializers.DB.First(&projection, id)
	return projection
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
