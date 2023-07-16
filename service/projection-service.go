package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type ProjectionService interface {
	Save(entity.Projection) entity.Projection
	FindAll() []entity.Projection
	FindOne(id string) entity.Projection
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
