package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type GenreService interface {
	Save(genre entity.Genre) error
	FindAll() []entity.Genre
	FindOne(id string) (entity.Genre, error)
	Update(id string, Genre entity.Genre) (entity.Genre, error)
	Delete(id string) error
}

type genreService struct {
	genres []entity.Genre
}

func NewGenreService() GenreService {
	return &genreService{}
}

func (service *genreService) Save(genre entity.Genre) error {
	err := initializers.DB.Create(&genre).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *genreService) FindOne(id string) (entity.Genre, error) {
	var genre entity.Genre
	result := initializers.DB.First(&genre, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return genre, errors.New("Genre not found")
		}
		return genre, result.Error
	}

	return genre, nil
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
