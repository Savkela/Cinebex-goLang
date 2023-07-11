package controller

import (
	"cinebex/entity"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	Save(ctx *gin.Context) entity.Movie
	FindAll() []entity.Movie
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

func (c *controller) Save(ctx *gin.Context) entity.Movie {
	var movie entity.Movie
	ctx.ShouldBindJSON(&movie)
	c.service.Save(movie)
	return movie
}
