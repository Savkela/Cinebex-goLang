package main

import (
	"cinebex/controller"
	"cinebex/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.MovieService       = service.New()
	videoController controller.MovieController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
