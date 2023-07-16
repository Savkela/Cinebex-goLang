package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeatController interface {
	Save(ctx *gin.Context) entity.Seat
	FindAll() []entity.Seat
	FindOne(ctx *gin.Context) entity.Seat
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type seatController struct {
	service service.SeatService
}

func NewSeatController(service service.SeatService) SeatController {
	return &seatController{
		service: service,
	}
}

func (c *seatController) FindAll() []entity.Seat {
	return c.service.FindAll()
}

func (c *seatController) FindOne(ctx *gin.Context) entity.Seat {
	id := ctx.Param("id")
	return c.service.FindOne(id)
}

func (c *seatController) Save(ctx *gin.Context) entity.Seat {
	var seat entity.Seat
	ctx.BindJSON(&seat)
	c.service.Save(seat)
	result := initializers.DB.Create(&seat)

	if result.Error != nil {
		ctx.Status(400)
		return seat
	}
	return seat
}

func (c *seatController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var seat entity.Seat
	if err := ctx.BindJSON(&seat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedSeat, err := c.service.Update(id, seat)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The seat is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedSeat)
}

func (c *seatController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Seat not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The Seat was successfully deleted"})
}
