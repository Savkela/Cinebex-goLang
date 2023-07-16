package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type RatingController interface {
	Save(ctx *gin.Context) entity.Rating
	FindAll() []entity.Rating
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

func (c *ratingController) Save(ctx *gin.Context) entity.Rating {
	var rating entity.Rating
	ctx.BindJSON(&rating)
	c.service.Save(rating)
	result := initializers.DB.Create(&rating)

	if result.Error != nil {
		ctx.Status(400)
		return rating
	}
	return rating
}
