package entity

type Movie struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Duration      string `json:"duration"`
	ImageUrl      string `json:"imageUrl"`
}
