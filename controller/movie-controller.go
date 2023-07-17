package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Movie
	FindOne(ctx *gin.Context)
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

func (c *controller) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *controller) Save(ctx *gin.Context) {
	var movie entity.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save movie"})
		return
	}

	ctx.JSON(http.StatusOK, movie)
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
