package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type MovieService interface {
	Save(entity.Movie) entity.Movie
	FindAll() []entity.Movie
	FindOne(id string) entity.Movie
	Update(id string, user entity.Movie) (entity.Movie, error)
	Delete(id string) error
}

type movieService struct {
	movies []entity.Movie
}

func New() MovieService {
	return &movieService{}
}

func (service *movieService) Save(movie entity.Movie) entity.Movie {
	service.movies = append(service.movies, movie)
	return movie
}

func (service *movieService) FindOne(id string) entity.Movie {
	var movie entity.Movie
	initializers.DB.First(&movie, id)
	return movie
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
