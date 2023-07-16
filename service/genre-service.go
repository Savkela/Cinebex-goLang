package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type GenreService interface {
	Save(entity.Genre) entity.Genre
	FindAll() []entity.Genre
	FindOne(id string) entity.Genre
	Update(id string, user entity.Genre) (entity.Genre, error)
	Delete(id string) error
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

func (service *genreService) Update(id string, genre entity.Genre) (entity.Genre, error) {
	var genreToUpdate entity.Genre
	err := initializers.DB.First(&genreToUpdate, id).Error
	if err != nil {
		return genreToUpdate, err
	}

	err = initializers.DB.Model(&genreToUpdate).Updates(genre).Error
	if err != nil {
		return genreToUpdate, err
	}

	return genreToUpdate, nil
}

func (service *genreService) Delete(id string) error {
	var genre entity.Genre
	result := initializers.DB.Delete(&genre, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
