package entity

type Movie struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Duration      string `json:"duration"`
	ImageUrl      string `json:"imageUrl"`
	GenreId       int64  `json:"genreId"`
}
