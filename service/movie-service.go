package service

import "cinebex/entity"

type MovieService interface {
	Save(entity.Movie) entity.Movie
	FindAll() []entity.Movie
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

func (service *movieService) FindAll() []entity.Movie {
	return service.movies
}

// func (service *movieService) Delete(movie entity.Movie) entity.Movie {
// 	service.movies = delete(movie)
// 	return movie
// }
