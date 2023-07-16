package entity

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Duration      string `json:"duration"`
	ImageUrl      string `json:"imageUrl"`
	GenreId       int64  `json:"genreId"`
}
