package controller

import (
	"cinebex/entity"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type ProjectionController interface {
	Save(ctx *gin.Context) entity.Projection
	FindAll() []entity.Projection
}

type projectionController struct {
	service service.ProjectionService
}

func NewProjectionController(service service.ProjectionService) ProjectionController {
	return &projectionController{
		service: service,
	}
}

func (c *projectionController) FindAll() []entity.Projection {
	return c.service.FindAll()
}

func (c *projectionController) Save(ctx *gin.Context) entity.Projection {
	var projection entity.Projection
	ctx.BindJSON(&projection)
	c.service.Save(projection)
	return projection
}
