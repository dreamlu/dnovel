package novel

import (
	"dnovel/models"
	"dnovel/models/novel"
	"dnovel/util/cons"
)

type BookService interface {
	GetClassify() []*models.Classify                                                 // 获得分类
	GetClassifyInfo(classify string) []*novel.BookInfo                               // 获得分类详情书本信息
	GetListByKeyword(keyword string) []*novel.BookInfo                               // 关键字搜索
	GetInfo(url string, source string) novel.BookInfo                                // 详情
	GetChapterList(url string, source string) []novel.Chapter                        // 章节列表
	GetContent(detailURL string, chapterURL string, source string) novel.BookContent // 章节内容
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

func (s *bookService) GetClassifyInfo(classify string) (itemList []*novel.BookInfo) {

	itemList = fService.GetClassifyInfo(classify)
	return
}

func (s *bookService) GetListByKeyword(keyword string) (itemList []*novel.BookInfo) {
	itemList = fService.GetItemList(keyword)
	return
}

func (s *bookService) GetInfo(url string, key string) (info novel.BookInfo) {
	info = fService.GetItem(url, key)
	return
}

func (s *bookService) GetChapterList(url string, key string) (chapterList []novel.Chapter) {
	chapterList = fService.GetChapterList(url, key)
	return
}

func (s *bookService) GetContent(detailURL string, chapterURL string, key string) (content novel.BookContent) {
	content = fService.GetContent(detailURL, chapterURL, key)
	return
}
