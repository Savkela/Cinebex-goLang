package controller

import (
	"cinebex/entity"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type ReservationController interface {
	Save(ctx *gin.Context) entity.Reservation
	FindAll() []entity.Reservation
}

type reservationController struct {
	service service.ReservationService
}

func NewReservationController(service service.ReservationService) ReservationController {
	return &reservationController{
		service: service,
	}
}

func (c *reservationController) FindAll() []entity.Reservation {
	return c.service.FindAll()
}

func (c *reservationController) Save(ctx *gin.Context) entity.Reservation {
	var reservation entity.Reservation
	ctx.BindJSON(&reservation)
	c.service.Save(reservation)
	return reservation
}
