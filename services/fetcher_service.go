package services

import (
	"dnovel/models/datamodels"
	"dnovel/util/fetcher"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type FetcherService interface {
	GetClassifyInfo(classify string) []*datamodels.BookInfo                                      // 分类书籍查找
	GetItemList(keyword string) []*datamodels.BookInfo                                           // 关键字查找
	GetItem(url string, key string) datamodels.BookInfo                                          // 书籍详情
	GetChapterList(url string, key string) []datamodels.Chapter                                  // 章节列表
	GetContent(detailURL string, chapterURL string, key string) (content datamodels.BookContent) // 获得内容
}

func NewFetcherService() FetcherService {
	return &fetcherService{}
}

var (
	sourceService = NewBookSourceService()
)

type fetcherService struct{}

// 分类书籍查找
func (s *fetcherService) GetClassifyInfo(classify string) (itemList []*datamodels.BookInfo) {
	bookSources := sourceService.GetAllSource()

	for i := range bookSources {
		source := &bookSources[i]

		for _, url := range source.ClassifyUrl[classify] {
			f := fetcher.NewFetcher()
			q, _ := queue.New(
				7, // Number of consumer threads
				&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
			)
			f.OnXML(source.ClassifyItemRule, func(e *colly.XMLElement) {
				itemList = append(itemList, s.parseClassifyInfo(source, e))
			})
			q.AddURL(url)
			q.Run(f)
		}
	}
	return
}

// 关键字搜索
func (s *fetcherService) GetItemList(keyword string) (itemList []*datamodels.BookInfo) {
	bookSources := sourceService.GetAllSource()

	for i := range bookSources {
		source := &bookSources[i]
		f := fetcher.NewFetcher()
		q, _ := queue.New(
			7, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
		)

		f.OnXML(source.SearchItemRule, func(e *colly.XMLElement) {
			itemList = append(itemList, s.parseItemSearch(source, e))
		})
		q.AddURL(fmt.Sprintf(source.SearchURL, keyword))
		q.Run(f)
	}
	return
}

func (s *fetcherService) GetItem(url string, key string) (info datamodels.BookInfo) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		return
	}

	f := fetcher.NewFetcher()
	f.OnXML(source.DetailBookItemRule, func(e *colly.XMLElement) {
		info = s.parseItemInfo(&source, e)
	})

	f.Visit(url)
	return
}

func (s *fetcherService) GetChapterList(url string, key string) (chapterList []datamodels.Chapter) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		return
	}

	//var chapterListURL string
	f := fetcher.NewFetcher()

	//if source.DetailChapterListURLRule != "" {
	//	f.OnXML(source.DetailChapterURLRule, func(e *colly.XMLElement) {
	//		var ele = fetcher.NewXMLElement(e)
	//		chapterListURL = ele.ChildUrl(source.DetailChapterListURLRule, "href")
	//		fc := fetcher.NewFetcher()
	//		fc.OnXML(source.DetailChapterRule, func(e *colly.XMLElement) {
	//			chapterList = append(chapterList, s.parseChapterList(&source, e, chapterListURL))
	//			return
	//		})
	//		fc.Visit(chapterListURL)
	//		return
	//	})
	//} else {
	f.OnXML(source.DetailChapterRule, func(e *colly.XMLElement) {
		chapterList = append(chapterList, s.parseChapterList(&source, e, url))
		return
	})
	//}
	f.Visit(url)
	return
}

func (s *fetcherService) GetContent(detailURL string, chapterURL string, key string) (content datamodels.BookContent) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		return
	}

	f := fetcher.NewFetcher()
	f.OnXML("//body", func(e *colly.XMLElement) {
		content = s.parseContent(&source, e, detailURL)
	})
	f.Visit(chapterURL)
	return
}

func (s *fetcherService) parseClassifyInfo(source *datamodels.BookSource, doc *colly.XMLElement) (item *datamodels.BookInfo) {
	var ele = fetcher.NewXMLElement(doc)
	item = &datamodels.BookInfo{
		Name:   ele.ChildText(source.ClassifyItemName),
		Author: ele.ChildText(source.ClassifyItemAuthor),
		URL:    ele.ChildUrl(source.ClassifyItemUrl, "href"),
		Source: source.SourceKey,
	}
	return
}

func (s *fetcherService) parseItemSearch(source *datamodels.BookSource, doc *colly.XMLElement) (item *datamodels.BookInfo) {
	var ele = fetcher.NewXMLElement(doc)
	item = &datamodels.BookInfo{
		Name:       ele.ChildText(source.SearchItemNameRule),
		Author:     ele.ChildText(source.SearchItemAuthorRule),
		Cover:      ele.ChildAttr(source.SearchItemCoverRule, "src"),
		NewChapter: ele.ChildText(source.SearchItemNewChapterRule),
		URL:        ele.ChildUrl(source.SearchItemURLRule, "href"),
		Source:     source.SourceKey,
	}
	return
}

func (s *fetcherService) parseItemInfo(source *datamodels.BookSource, doc *colly.XMLElement) (info datamodels.BookInfo) {
	var ele = fetcher.NewXMLElement(doc)
	info = datamodels.BookInfo{
		Name:        ele.ChildText(source.DetailBookNameRule),
		Author:      ele.ChildText(source.DetailBookAuthorRule),
		Cover:       ele.ChildAttr(source.DetailBookCoverRule, "src"),
		Category:    ele.ChildText(source.DetailBookCategoryRule),
		Description: ele.ChildHtml(source.DetailBookDescriptionRule),
		NewChapter:  ele.ChildText(source.DetailBookNewChapterRule),
		URL:         ele.ChildUrl(source.DetailBookNewChapterUrlRule, "href"),
		Source:      source.SourceKey,
	}
	// 这里应该用爬虫自动处理才对,没找到方法~呜呜~,(找到了啦啦啦~ChildUrl)
	//if !strings.Contains(info.URL, "http") {
	//	url := doc.Request.URL
	//	if !strings.Contains(info.URL, "/") {
	//		info.URL = url.String() + info.URL
	//	} else {
	//		info.URL = url.Scheme + "://" + url.Host + info.URL
	//	}
	//}
	return
}

func (s *fetcherService) parseChapterList(source *datamodels.BookSource, doc *colly.XMLElement, url string) (chapter datamodels.Chapter) {
	var ele = fetcher.NewXMLElement(doc)
	chapter = datamodels.Chapter{
		Title:      ele.ChildText(source.DetailChapterTitleRule),
		DetailURL:  url,
		ChapterURL: ele.ChildUrl(source.DetailChapterURLRule, "href"),
		Source:     source.SourceKey,
	}

	return
}

func (s *fetcherService) parseContent(source *datamodels.BookSource, doc *colly.XMLElement, url string) (content datamodels.BookContent) {
	var ele = fetcher.NewXMLElement(doc)
	content = datamodels.BookContent{
		Title:       ele.ChildText(source.ContentTitleRule),
		Text:        ele.ChildHtml(source.ContentTextRule),
		DetailURL:   url,
		PreviousURL: ele.ChildUrl(source.ContentPreviousURLRule, "href"),
		NextURL:     ele.ChildUrl(source.ContentNextURLRule, "href"),
		Source:      source.SourceKey,
	}
	return content
}
