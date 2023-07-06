package service

import "cinebex/entity"

type GenreService interface {
	Save(entity.Genre) entity.Genre
	FindAll() []entity.Genre
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

func (service *genreService) FindAll() []entity.Genre {
	return service.genres
}
