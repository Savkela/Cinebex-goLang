package controller

import (
	"cinebex/entity"
	"cinebex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Reservation
	FindOne(ctx *gin.Context)
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

func (c *reservationController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	reservation, err := c.service.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

func (c *reservationController) Save(ctx *gin.Context) {
	var reservation entity.Reservation
	if err := ctx.BindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.Save(reservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reservation"})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
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
