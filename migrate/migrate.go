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
	initializers.DB.AutoMigrate(&entity.User{})
}
