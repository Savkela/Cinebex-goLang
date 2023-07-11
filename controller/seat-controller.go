package controller

import (
	"cinebex/entity"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type SeatController interface {
	Save(ctx *gin.Context) entity.Seat
	FindAll() []entity.Seat
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

func (c *seatController) Save(ctx *gin.Context) entity.Seat {
	var seat entity.Seat
	ctx.BindJSON(&seat)
	c.service.Save(seat)
	return seat
}
