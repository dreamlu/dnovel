package novel

import (
	novel2 "dnovel/datasource/novel"
	"dnovel/models/novel"
	"dnovel/util/repositories"
)

type BookSourceService interface {
	GetSourceByKey(key string) (novel.BookSource, bool)
	GetAllSource() []novel.BookSource
}

func NewBookSourceService() BookSourceService {
	return &bookSourceService{
		repositories.NewBookSourceRepository(novel2.BookSources),
	}
}

type bookSourceService struct {
	rep repositories.BookSourceRepository
}

func (s *bookSourceService) GetSourceByKey(key string) (novel.BookSource, bool) {
	return s.rep.Select(func(s novel.BookSource) bool {
		return s.SourceKey == key
	})
}

func (s *bookSourceService) GetAllSource() []novel.BookSource {
	return s.rep.SelectMany(func(_ novel.BookSource) bool {
		return true
	}, -1)
}
