package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type MovieService interface {
	Save(movie entity.Movie) error
	FindAll() []entity.Movie
	FindOne(id string) (entity.Movie, error)
	Update(id string, Movie entity.Movie) (entity.Movie, error)
	Delete(id string) error
}

type movieService struct {
	movies []entity.Movie
}

func New() MovieService {
	return &movieService{}
}

func (service *movieService) Save(movie entity.Movie) error {
	err := initializers.DB.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *movieService) FindOne(id string) (entity.Movie, error) {
	var movie entity.Movie
	result := initializers.DB.First(&movie, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return movie, errors.New("Movie not found")
		}
		return movie, result.Error
	}

	return movie, nil
}
func (service *movieService) FindAll() []entity.Movie {
	var movies []entity.Movie
	initializers.DB.Find(&movies)
	return movies
}

func (service *movieService) Update(id string, movie entity.Movie) (entity.Movie, error) {
	var movieToUpdate entity.Movie
	err := initializers.DB.First(&movieToUpdate, id).Error
	if err != nil {
		return movieToUpdate, err
	}

	err = initializers.DB.Model(&movieToUpdate).Updates(movie).Error
	if err != nil {
		return movieToUpdate, err
	}

	return movieToUpdate, nil
}

func (service *movieService) Delete(id string) error {
	var movie entity.Movie
	result := initializers.DB.Delete(&movie, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
