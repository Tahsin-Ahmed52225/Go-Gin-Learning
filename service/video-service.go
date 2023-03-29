package service

import "./Go-Gin-Learning/entity"

type VideoService struct {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video


}
