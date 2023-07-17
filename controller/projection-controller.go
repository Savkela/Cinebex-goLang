package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectionController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Projection
	FindOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func (c *projectionController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	projection, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, projection)
}

func (c *projectionController) Save(ctx *gin.Context) {
	var projection entity.Projection
	if err := ctx.BindJSON(&projection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(projection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save projection"})
		return
	}

	ctx.JSON(http.StatusOK, projection)
}

func (c *projectionController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var projection entity.Projection
	if err := ctx.BindJSON(&projection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedProjection, err := c.service.Update(id, projection)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The Projection is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedProjection)
}

func (c *projectionController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Projection not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The Projection was successfully deleted"})
}
