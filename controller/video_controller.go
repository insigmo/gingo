package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/insigmo/gingo/entity"
	"github.com/insigmo/gingo/service"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &controller{service: s}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	c.service.Save(video)

	if err != nil {
		panic(err)
	}
	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
