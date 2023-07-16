package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	var users []entity.User
	initializers.DB.Find(&users)
	return users
}
