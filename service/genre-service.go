package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type GenreService interface {
	Save(entity.Genre) entity.Genre
	FindAll() []entity.Genre
	FindOne(id string) entity.Genre
}

type genreService struct {
	genres []entity.Genre
}

func NewGenreService() GenreService {
	return &genreService{}
}

func (service *genreService) Save(genre entity.Genre) entity.Genre {
	service.genres = append(service.genres, genre)
	return genre
}

func (service *genreService) FindOne(id string) entity.Genre {
	var genre entity.Genre
	initializers.DB.First(&genre, id)
	return genre
}

func (service *genreService) FindAll() []entity.Genre {
	var genres []entity.Genre
	initializers.DB.Find(&genres)
	return genres
}
