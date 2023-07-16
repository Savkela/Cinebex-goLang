package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type ProjectionController interface {
	Save(ctx *gin.Context) entity.Projection
	FindAll() []entity.Projection
	FindOne(ctx *gin.Context) entity.Projection
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

func (c *projectionController) FindOne(ctx *gin.Context) entity.Projection {
	id := ctx.Param("id")
	return c.service.FindOne(id)
}

func (c *projectionController) Save(ctx *gin.Context) entity.Projection {
	var projection entity.Projection
	ctx.BindJSON(&projection)
	c.service.Save(projection)
	result := initializers.DB.Create(&projection)

	if result.Error != nil {
		ctx.Status(400)
		return projection
	}
	return projection
}
