package controller

import (
	"cinebex/entity"
	"cinebex/initializers"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Save(ctx *gin.Context) entity.User
	FindAll() []entity.User
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.Status(400)
		return user
	}
	return user
}
