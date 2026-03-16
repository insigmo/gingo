package service

import "github.com/insigmo/gingo/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (v *videoService) Save(video entity.Video) entity.Video {
	v.videos = append(v.videos, video)
	return video
}

func (v *videoService) FindAll() []entity.Video {
	return v.videos
}
