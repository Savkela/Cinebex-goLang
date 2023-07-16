package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type MovieService interface {
	Save(entity.Movie) entity.Movie
	FindAll() []entity.Movie
	FindOne(id string) entity.Movie
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
