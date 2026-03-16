// Gin project
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/insigmo/gingo/controller"
	"github.com/insigmo/gingo/middlewares"
	"github.com/insigmo/gingo/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
