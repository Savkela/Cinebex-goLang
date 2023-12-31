package service

import "cinebex/entity"

type ProjectionService interface {
	Save(entity.Projection) entity.Projection
	FindAll() []entity.Projection
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

func (service *projectionService) FindAll() []entity.Projection {
	return service.projections
}
