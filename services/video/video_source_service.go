package video

import (
	video2 "dnovel/datasource/video"
	"dnovel/models/video"
)

type VideoSourceService interface {
	GetSourceByKey(key string) video.VideoSource
	GetAllSource() map[string]video.VideoSource
}

func NewVideoSourceService() VideoSourceService {
	return &bookSourceService{
		video2.VideoSources,
	}
}

type bookSourceService struct {
	source map[string]video.VideoSource
}

func (s *bookSourceService) GetSourceByKey(key string) video.VideoSource {
	return s.source[key]
}

func (s *bookSourceService) GetAllSource() map[string]video.VideoSource {
	return s.source
}
