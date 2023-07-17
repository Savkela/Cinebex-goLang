package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatingController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Rating
	FindOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ratingController struct {
	service service.RatingService
}

func NewRatingController(service service.RatingService) RatingController {
	return &ratingController{
		service: service,
	}
}

func (c *ratingController) FindAll() []entity.Rating {
	return c.service.FindAll()
}

func (c *ratingController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	rating, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rating)
}

func (c *ratingController) Save(ctx *gin.Context) {
	var rating entity.Rating
	if err := ctx.BindJSON(&rating); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save rating"})
		return
	}

	ctx.JSON(http.StatusOK, rating)
}

func (c *ratingController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var rating entity.Rating
	if err := ctx.BindJSON(&rating); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedRating, err := c.service.Update(id, rating)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The rating is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedRating)
}

func (c *ratingController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The rating was successfully deleted"})
}
