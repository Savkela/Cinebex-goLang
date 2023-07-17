package main

import (
	"cinebex/controller"
	"cinebex/initializers"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

var (
	movieService    service.MovieService       = service.New()
	movieController controller.MovieController = controller.New(movieService)

	projectionService    service.ProjectionService       = service.NewProjectionService()
	projectionController controller.ProjectionController = controller.NewProjectionController(projectionService)

	userService    service.UserService       = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService)

	genreService    service.GenreService       = service.NewGenreService()
	genreController controller.GenreController = controller.NewGenreController(genreService)

	ratingService    service.RatingService       = service.NewRatingService()
	ratingController controller.RatingController = controller.NewRatingController(ratingService)

	reservationService    service.ReservationService       = service.NewReservationService()
	reservationController controller.ReservationController = controller.NewReservationController(reservationService)

	seatService    service.SeatService       = service.NewSeatService()
	seatController controller.SeatController = controller.NewSeatController(seatService)
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), gin.BasicAuth(gin.Accounts{"savic": "nikola"}))

	//movies
	server.GET("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, movieController.FindAll())
	})

	server.POST("/movies", movieController.Save)

	server.GET("/movies/:id", movieController.FindOne)

	server.PUT("/movies/:id", movieController.Update)

	server.DELETE("/movies/:id", movieController.Delete)

	//projections
	server.GET("/projections", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.FindAll())
	})

	server.POST("/projections", projectionController.Save)

	server.GET("/projections/:id", projectionController.FindOne)

	server.PUT("/projections/:id", projectionController.Update)

	server.DELETE("/projections/:id", projectionController.Delete)

	//users
	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", userController.Save)

	server.GET("/users/:id", userController.FindOne)

	server.PUT("/users/:id", userController.Update)

	server.DELETE("/users/:id", userController.Delete)

	//genres
	server.GET("/genres", func(ctx *gin.Context) {
		ctx.JSON(200, genreController.FindAll())
	})

	server.POST("/genres", genreController.Save)

	server.GET("/genres/:id", genreController.FindOne)

	server.PUT("/genres/:id", genreController.Update)

	server.DELETE("/genres/:id", genreController.Delete)

	//reservations
	server.GET("/reservations", func(ctx *gin.Context) {
		ctx.JSON(200, reservationController.FindAll())
	})

	server.POST("/reservations", reservationController.Save)

	server.GET("/reservations/:id", reservationController.FindOne)

	server.PUT("/reservations/:id", reservationController.Update)

	server.DELETE("/reservations/:id", reservationController.Delete)

	//seats
	server.GET("/seats", func(ctx *gin.Context) {
		ctx.JSON(200, seatController.FindAll())
	})

	server.POST("/seats", seatController.Save)

	server.GET("/seats/:id", seatController.FindOne)

	server.PUT("/seats/:id", seatController.Update)

	server.DELETE("/seats/:id", seatController.Delete)

	//ratings
	server.GET("/ratings", func(ctx *gin.Context) {
		ctx.JSON(200, ratingController.FindAll())
	})

	server.POST("/ratings", ratingController.Save)

	server.GET("/ratings/:id", ratingController.FindOne)

	server.PUT("/ratings/:id", ratingController.Update)

	server.DELETE("/ratings/:id", ratingController.Delete)

	server.Run()
}
