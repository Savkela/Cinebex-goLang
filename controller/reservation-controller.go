package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController interface {
	Save(ctx *gin.Context) entity.Reservation
	FindAll() []entity.Reservation
	FindOne(ctx *gin.Context) entity.Reservation
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func (c *reservationController) FindOne(ctx *gin.Context) entity.Reservation {
	id := ctx.Param("id")
	return c.service.FindOne(id)
}

func (c *reservationController) Save(ctx *gin.Context) entity.Reservation {
	var reservation entity.Reservation
	ctx.BindJSON(&reservation)
	c.service.Save(reservation)
	result := initializers.DB.Create(&reservation)

	if result.Error != nil {
		ctx.Status(400)
		return reservation
	}
	return reservation
}

func (c *reservationController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var reservation entity.Reservation
	if err := ctx.BindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nonsense JSON request"})
		return
	}

	updatedReservation, err := c.service.Update(id, reservation)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "The reservation is not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedReservation)
}

func (c *reservationController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found or error deleting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The Reservation was successfully deleted"})
}
