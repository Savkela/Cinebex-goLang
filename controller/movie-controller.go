package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	Save(ctx *gin.Context) entity.Movie
	FindAll() []entity.Movie
	FindOne(ctx *gin.Context) entity.Movie
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func (c *controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var movie entity.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedMovie, err := c.service.Update(id, movie)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The Movie is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedMovie)
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The Movie was successfully deleted"})
}
