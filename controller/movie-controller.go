package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	Save(ctx *gin.Context) entity.Movie
	FindAll() []entity.Movie
	FindOne(ctx *gin.Context) entity.Movie
}

type controller struct {
	service service.MovieService
}

func New(service service.MovieService) MovieController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Movie {
	return c.service.FindAll()
}

func (c *controller) FindOne(ctx *gin.Context) entity.Movie {
	id := ctx.Param("id")
	return c.service.FindOne(id)
}

func (c *controller) Save(ctx *gin.Context) entity.Movie {
	var movie entity.Movie
	ctx.ShouldBindJSON(&movie)
	c.service.Save(movie)
	result := initializers.DB.Create(&movie)

	if result.Error != nil {
		ctx.Status(400)
		return movie
	}
	return movie
}
