package services

import (
	"dnovel/models/datamodels"
	"dnovel/util/cons"
)

type BookService interface {
	GetClassify() []*datamodels.Classify                             // 获得分类
	GetClassifyInfo(classify string) []*datamodels.BookInfo          // 获得分类详情书本信息
	GetSearch(keyword string) []*datamodels.BookInfo                 // 关键字搜索
	GetInfo(url string, source string) datamodels.BookInfo           // 详情
	GetChapterList(url string, source string) []datamodels.Chapter   // 章节列表
	GetRead(chapterURL string, source string) datamodels.BookContent // 章节内容
}

func NewBookService() BookService {
	return &bookService{}
}

type bookService struct{}

var (
	fService = NewFetcherService()
)

var Classifys = []string{cons.Classify1, cons.Classify2, cons.Classify3, cons.Classify4, cons.Classify5, cons.Classify6}

func (s *bookService) GetClassify() (cs []*datamodels.Classify) {

	for _, v := range Classifys {
		cs = append(cs, &datamodels.Classify{Name: v})
	}
	return
}

func (s *bookService) GetClassifyInfo(classify string) (itemList []*datamodels.BookInfo) {

	itemList = fService.GetClassifyInfo(classify)
	return
}

func (s *bookService) GetSearch(keyword string) (itemList []*datamodels.BookInfo) {
	itemList = fService.GetSearch(keyword)
	return
}

func (s *bookService) GetInfo(url string, key string) (info datamodels.BookInfo) {
	info = fService.GetInfo(url, key)
	return
}

func (s *bookService) GetChapterList(url string, key string) (chapterList []datamodels.Chapter) {
	chapterList = fService.GetChapterList(url, key)
	return
}

func (s *bookService) GetRead(chapterURL string, key string) (content datamodels.BookContent) {
	content = fService.GetRead(chapterURL, key)
	return
}
