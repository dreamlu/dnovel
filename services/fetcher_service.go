package services

import (
	"dnovel/models/datamodels"
	"dnovel/util/fetcher"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"strings"
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
				10, // Number of consumer threads
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

		url := fmt.Sprintf(source.SearchURL, keyword)
		q, _ := queue.New(
			10, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
		)
		// 反爬破解
		//f.OnResponse(func(res *colly.Response) {
		//	if source.SourceKey == "biquge" {
		//		newUrl := resBody(res.Body)
		//		if newUrl != "" {
		//			q, _ := queue.New(
		//				10, // Number of consumer threads
		//				&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
		//			)
		//			q.AddURL(url + newUrl)
		//			q.Run(f)
		//		}
		//	}
		//})

		f.OnXML(source.SearchItemRule, func(e *colly.XMLElement) {
			itemList = append(itemList, s.parseItemSearch(source, e))
		})
		q.AddURL(url)
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

	url = strings.Replace(url, "m.", "www.", 1)
	f.Visit(url)
	return
}

func (s *fetcherService) GetChapterList(url string, key string) (chapterList []datamodels.Chapter) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		return
	}

	f := fetcher.NewFetcher()
	//f.Async = true // f.wait() // 异步执行
	q, _ := queue.New(
		10, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 100}, // Use default queue storage
	)

	f.OnXML(source.DetailChapterRule, func(e *colly.XMLElement) {
		chapterList = append(chapterList, s.parseChapterList(&source, e, url))
		return
	})
	url = strings.Replace(url, "m.", "www.", 1)
	q.AddURL(url)
	q.Run(f)
	return
}

func (s *fetcherService) GetContent(detailURL string, chapterURL string, key string) (content datamodels.BookContent) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		return
	}

	f := fetcher.NewFetcher()
	f.OnXML("//body", func(e *colly.XMLElement) {
		content = s.parseContent(&source, e, chapterURL)
	})
	f.Visit(chapterURL)
	return
}

func (s *fetcherService) parseClassifyInfo(source *datamodels.BookSource, doc *colly.XMLElement) (item *datamodels.BookInfo) {
	var (
		ele   = fetcher.NewXMLElement(doc)
		cover = "src"
	)
	if source.SourceKey == "beqegecc" {
		cover = "data-original"
	}
	item = &datamodels.BookInfo{
		Name:        ele.ChildText(source.ClassifyItemName),
		Author:      ele.ChildText(source.ClassifyItemAuthor),
		URL:         ele.ChildUrl(source.ClassifyItemUrl, "href"),
		Cover:       ele.ChildUrl(source.ClassifyItemCover, cover),
		Description: ele.ChildText(source.ClassifyItemDesc),
		Source:      source.SourceKey,
	}
	item.URL = strings.Replace(item.URL, "m.", "www.", 1)
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
	item.URL = strings.Replace(item.URL, "m.", "", 1)
	if strings.Contains(item.Author, "：") {
		item.Author = strings.Split(item.Author, "：")[1]
	}
	return
}

var (
	auts = []string{":", "："}
)

func (s *fetcherService) parseItemInfo(source *datamodels.BookSource, doc *colly.XMLElement) (info datamodels.BookInfo) {
	var ele = fetcher.NewXMLElement(doc)
	info = datamodels.BookInfo{
		Name:         ele.ChildText(source.DetailBookNameRule),
		Author:       ele.ChildText(source.DetailBookAuthorRule),
		Cover:        ele.ChildUrlText(source.DetailBookCoverRule),
		Category:     ele.ChildText(source.DetailBookCategoryRule),
		Description:  ele.ChildHtml(source.DetailBookDescriptionRule),
		NewChapter:   ele.ChildText(source.DetailBookNewChapterRule),
		URL:          ele.ChildUrl(source.DetailBookNewChapterUrlRule, "href"),
		FirstChapter: ele.ChildText(source.DetailBookFirstChapterRule),
		FirstURL:     ele.ChildUrl(source.DetailBookFirstUrlRule, "href"),
		Source:       source.SourceKey,
	}
	for _, v := range auts {
		if strings.Contains(info.Author, v) {
			info.Author = strings.Split(info.Author, v)[1]
		}
	}
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
		Text:        ele.ChildRemoveHtml(source.ContentTextRule, "div"),
		CurrentURL:  url,
		PreviousURL: ele.ChildUrl(source.ContentPreviousURLRule, "href"),
		NextURL:     ele.ChildUrl(source.ContentNextURLRule, "href"),
		Source:      source.SourceKey,
	}
	// bequgee
	switch source.SourceKey {
	case "ibiquge":
		i := strings.Index(content.Text, "</p>")
		if i != -1 {
			content.Text = content.Text[i+4:]
		}
	case "biqugee":
		i := strings.Index(content.Text, "<p><a")
		k := strings.Index(content.Text, "a></p>")
		if i != -1 {
			content.Text = content.Text[:i] + content.Text[k+6:]
		}
	}
	return content
}
