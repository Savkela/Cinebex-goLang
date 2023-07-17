package service

import (
	"cinebex/entity"
	"cinebex/initializers"
	"errors"

	"gorm.io/gorm"
)

type RatingService interface {
	Save(rating entity.Rating) error
	FindOne(id string) (entity.Rating, error)
	FindAll() []entity.Rating
	Update(id string, rating entity.Rating) (entity.Rating, error)
	Delete(id string) error
}

type ratingService struct {
	ratings []entity.Rating
}

func NewRatingService() RatingService {
	return &ratingService{}
}

func (service *ratingService) Save(rating entity.Rating) error {
	err := initializers.DB.Create(&rating).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *ratingService) FindOne(id string) (entity.Rating, error) {
	var rating entity.Rating
	result := initializers.DB.First(&rating, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return rating, errors.New("Rating not found")
		}
		return rating, result.Error
	}

	return rating, nil
}

func (service *ratingService) FindAll() []entity.Rating {
	var ratings []entity.Rating
	initializers.DB.Find(&ratings)
	return ratings
}

func (service *ratingService) Update(id string, rating entity.Rating) (entity.Rating, error) {
	var ratingToUpdate entity.Rating
	err := initializers.DB.First(&ratingToUpdate, id).Error
	if err != nil {
		return ratingToUpdate, err
	}

	err = initializers.DB.Model(&ratingToUpdate).Updates(rating).Error
	if err != nil {
		return ratingToUpdate, err
	}

	return ratingToUpdate, nil
}

func (service *ratingService) Delete(id string) error {
	var rating entity.Rating
	result := initializers.DB.Delete(&rating, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
