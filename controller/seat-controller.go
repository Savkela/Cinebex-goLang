package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeatController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Seat
	FindOne(ctx *gin.Context)
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

func (c *seatController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	seat, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seat)
}

func (c *seatController) Save(ctx *gin.Context) {
	var seat entity.Seat
	if err := ctx.BindJSON(&seat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(seat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save seat"})
		return
	}

	ctx.JSON(http.StatusOK, seat)
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
