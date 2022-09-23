package video

import (
	"dnovel/models"
	"dnovel/models/video"
	"dnovel/util/cons"
)

type BookService interface {
	GetClassify() []*models.Classify                                                  // 获得分类
	GetClassifyInfo(classify string) []*video.VideoInfo                               // 获得分类详情书本信息
	GetListByKeyword(keyword string) []*video.VideoInfo                               // 关键字搜索
	GetInfo(url string, source string) video.VideoInfo                                // 详情
	GetChapterList(url string, source string) []video.Chapter                         // 章节列表
	GetContent(detailURL string, chapterURL string, source string) video.VideoContent // 章节内容
}

func NewBookService() BookService {
	return &bookService{}
}

type bookService struct{}

var (
	fService = NewFetcherService()
)

var Classifys = []string{cons.Classify1, cons.Classify2, cons.Classify3, cons.Classify4, cons.Classify5, cons.Classify6}

func (s *bookService) GetClassify() (cs []*models.Classify) {

	for _, v := range Classifys {
		cs = append(cs, &models.Classify{Name: v})
	}
	return
}

func (s *bookService) GetClassifyInfo(classify string) (itemList []*video.VideoInfo) {

	itemList = fService.GetClassifyInfo(classify)
	return
}

func (s *bookService) GetListByKeyword(keyword string) (itemList []*video.VideoInfo) {
	itemList = fService.GetItemList(keyword)
	return
}

func (s *bookService) GetInfo(url string, key string) (info video.VideoInfo) {
	info = fService.GetItem(url, key)
	return
}

func (s *bookService) GetChapterList(url string, key string) (chapterList []video.Chapter) {
	chapterList = fService.GetChapterList(url, key)
	return
}

func (s *bookService) GetContent(detailURL string, chapterURL string, key string) (content video.VideoContent) {
	content = fService.GetContent(detailURL, chapterURL, key)
	return
}
