package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	FindOne(id string) entity.User
	Update(id string, user entity.User) (entity.User, error)
	Delete(id string) error
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

func (service *userService) FindOne(id string) entity.User {
	var user entity.User
	initializers.DB.First(&user, id)
	return user
}

func (service *userService) FindAll() []entity.User {
	var users []entity.User
	initializers.DB.Find(&users)
	return users
}

func (service *userService) Update(id string, user entity.User) (entity.User, error) {
	var userToUpdate entity.User
	err := initializers.DB.First(&userToUpdate, id).Error
	if err != nil {
		return userToUpdate, err
	}

	err = initializers.DB.Model(&userToUpdate).Updates(user).Error
	if err != nil {
		return userToUpdate, err
	}

	return userToUpdate, nil
}

func (service *userService) Delete(id string) error {
	var user entity.User
	result := initializers.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
