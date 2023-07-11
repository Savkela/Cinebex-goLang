package controller

import (
	"cinebex/entity"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type GenreController interface {
	Save(ctx *gin.Context) entity.Genre
	FindAll() []entity.Genre
}

type genreController struct {
	service service.GenreService
}

func NewGenreController(service service.GenreService) GenreController {
	return &genreController{
		service: service,
	}
}

func (c *genreController) FindAll() []entity.Genre {
	return c.service.FindAll()
}

func (c *genreController) Save(ctx *gin.Context) entity.Genre {
	var genre entity.Genre
	ctx.BindJSON(&genre)
	c.service.Save(genre)
	return genre
}
