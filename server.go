package main

import (
	"cinebex/controller"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

var (
	movieService    service.MovieService       = service.New()
	movieController controller.MovieController = controller.New(movieService)

	projectionService    service.ProjectionService       = service.NewProjectionService()
	projectionController controller.ProjectionController = controller.NewProjectionController(projectionService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), gin.BasicAuth(gin.Accounts{"savic": "nikola"}))

	server.GET("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, movieController.FindAll())
	})

	server.POST("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, movieController.Save(ctx))
	})

	server.GET("/projections", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.FindAll())
	})

	server.POST("/projections", func(ctx *gin.Context) {
		ctx.JSON(200, projectionController.Save(ctx))

	})

	server.Run(":8080")
}
