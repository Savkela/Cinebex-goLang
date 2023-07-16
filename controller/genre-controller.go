package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreController interface {
	Save(ctx *gin.Context) entity.Genre
	FindAll() []entity.Genre
	FindOne(ctx *gin.Context) entity.Genre
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

func (c *genreController) FindOne(ctx *gin.Context) entity.Genre {
	id := ctx.Param("id")
	return c.service.FindOne(id)
}

func (c *genreController) Save(ctx *gin.Context) entity.Genre {
	var genre entity.Genre
	ctx.BindJSON(&genre)
	c.service.Save(genre)
	result := initializers.DB.Create(&genre)

	if result.Error != nil {
		ctx.Status(400)
		return genre
	}
	return genre
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
