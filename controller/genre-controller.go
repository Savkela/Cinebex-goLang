package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Genre
	FindOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func (c *genreController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	genre, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, genre)
}

func (c *genreController) Save(ctx *gin.Context) {
	var genre entity.Genre
	if err := ctx.BindJSON(&genre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(genre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save genre"})
		return
	}

	ctx.JSON(http.StatusOK, genre)
}

func (c *genreController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var genre entity.Genre
	if err := ctx.BindJSON(&genre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedGenre, err := c.service.Update(id, genre)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The Genre is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedGenre)
}

func (c *genreController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Genre not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The Genre was successfully deleted"})
}
