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

	server.POST("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, movieController.Save(ctx))
	})

	server.GET("/movies/:id", func(ctx *gin.Context) {
		ctx.JSON(200, movieController.FindOne(ctx))
	})

	server.PUT("/movies/:id", movieController.Update)

	server.DELETE("/movies/:id", movieController.Delete)

	//projections
	server.GET("/projections", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.FindAll())
	})

	server.POST("/projections", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.Save(ctx))
	})

	server.GET("/projections/:id", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.FindOne(ctx))
	})

	server.PUT("/projections/:id", projectionController.Update)

	server.DELETE("/projections/:id", projectionController.Delete)

	//users
	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	server.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindOne(ctx))
	})

	server.PUT("/users/:id", userController.Update)

	server.DELETE("/users/:id", userController.Delete)

	//genres
	server.GET("/genres", func(ctx *gin.Context) {
		ctx.JSON(200, genreController.FindAll())
	})

	server.POST("/genres", func(ctx *gin.Context) {
		ctx.JSON(200, genreController.Save(ctx))
	})

	server.GET("/genres/:id", func(ctx *gin.Context) {
		ctx.JSON(200, genreController.FindOne(ctx))
	})

	server.PUT("/genres/:id", genreController.Update)

	server.DELETE("/genres/:id", genreController.Delete)

	//reservations
	server.GET("/reservations", func(ctx *gin.Context) {
		ctx.JSON(200, reservationController.FindAll())
	})

	server.POST("/reservations", func(ctx *gin.Context) {
		ctx.JSON(200, reservationController.Save(ctx))
	})

	server.GET("/reservations/:id", func(ctx *gin.Context) {
		ctx.JSON(200, reservationController.FindOne(ctx))
	})

	server.PUT("/reservations/:id", reservationController.Update)

	server.DELETE("/reservations/:id", reservationController.Delete)

	//seats
	server.GET("/seats", func(ctx *gin.Context) {
		ctx.JSON(200, seatController.FindAll())
	})

	server.POST("/seats", func(ctx *gin.Context) {
		ctx.JSON(200, seatController.Save(ctx))
	})

	server.GET("/seats/:id", func(ctx *gin.Context) {
		ctx.JSON(200, seatController.FindOne(ctx))
	})

	server.PUT("/seats/:id", seatController.Update)

	server.DELETE("/seats/:id", seatController.Delete)

	//ratings
	server.GET("/ratings", func(ctx *gin.Context) {
		ctx.JSON(200, ratingController.FindAll())
	})

	server.POST("/ratings", func(ctx *gin.Context) {
		ctx.JSON(200, ratingController.Save(ctx))
	})

	server.GET("/ratings/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ratingController.FindOne(ctx))
	})

	server.PUT("/ratings/:id", ratingController.Update)

	server.DELETE("/ratings/:id", ratingController.Delete)

	server.Run()
}
