package main

import (
	"cinebex/entity"
	"cinebex/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&entity.User{}, &entity.Genre{}, &entity.Movie{}, &entity.Projection{}, &entity.Rating{}, &entity.Reservation{}, &entity.Seat{}, &entity.Genre{})
}
