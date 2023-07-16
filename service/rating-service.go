package service

import (
	"cinebex/entity"
	"cinebex/initializers"
)

type RatingService interface {
	Save(entity.Rating) entity.Rating
	FindAll() []entity.Rating
}

type ratingService struct {
	ratings []entity.Rating
}

func NewRatingService() RatingService {
	return &ratingService{}
}

func (service *ratingService) Save(rating entity.Rating) entity.Rating {
	service.ratings = append(service.ratings, rating)
	return rating
}

func (service *ratingService) FindAll() []entity.Rating {
	var ratings []entity.Rating
	initializers.DB.Find(&ratings)
	return ratings
}
