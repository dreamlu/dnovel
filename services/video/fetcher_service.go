package video

import (
	"dnovel/models/video"
	"dnovel/util/fetcher"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"strings"
)

type FetcherService interface {
	GetClassifyInfo(classify string) []*video.VideoInfo                                      // 分类书籍查找
	GetItemList(keyword string) []*video.VideoInfo                                           // 关键字查找
	GetItem(url string, key string) video.VideoInfo                                          // 书籍详情
	GetChapterList(url string, key string) []video.Chapter                                   // 章节列表
	GetContent(detailURL string, chapterURL string, key string) (content video.VideoContent) // 获得内容
}

func NewFetcherService() FetcherService {
	return &fetcherService{}
}

var (
	sourceService = NewVideoSourceService()
)

type fetcherService struct{}

// 分类书籍查找
func (s *fetcherService) GetClassifyInfo(classify string) (itemList []*video.VideoInfo) {
	bookSources := sourceService.GetAllSource()

	for i := range bookSources {
		source := bookSources[i]

		for _, url := range source.ClassifyUrl[classify] {
			f := fetcher.NewFetcher()
			q, _ := queue.New(
				10, // Number of consumer threads
				&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
			)
			f.OnXML(source.ClassifyItemRule, func(e *colly.XMLElement) {
				itemList = append(itemList, s.parseClassifyInfo(&source, e))
			})
			q.AddURL(url)
			q.Run(f)
		}
	}
	return
}

// 关键字搜索
func (s *fetcherService) GetItemList(keyword string) (itemList []*video.VideoInfo) {
	bookSources := sourceService.GetAllSource()

	for i := range bookSources {
		source := bookSources[i]
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
			itemList = append(itemList, s.parseItemSearch(&source, e))
		})
		q.AddURL(url)
		q.Run(f)
	}
	return
}

func (s *fetcherService) GetItem(url string, key string) (info video.VideoInfo) {
	source := sourceService.GetSourceByKey(key)

	f := fetcher.NewFetcher()
	f.OnXML(source.DetailVideoItemRule, func(e *colly.XMLElement) {
		info = s.parseItemInfo(&source, e)
	})

	url = strings.Replace(url, "m.", "www.", 1)
	f.Visit(url)
	return
}

func (s *fetcherService) GetChapterList(url string, key string) (chapterList []video.Chapter) {
	source := sourceService.GetSourceByKey(key)

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

func (s *fetcherService) GetContent(detailURL string, chapterURL string, key string) (content video.VideoContent) {
	source := sourceService.GetSourceByKey(key)

	f := fetcher.NewFetcher()
	f.OnXML("//body", func(e *colly.XMLElement) {
		content = s.parseContent(&source, e, chapterURL)
	})
	f.Visit(chapterURL)
	return
}

func (s *fetcherService) parseClassifyInfo(source *video.VideoSource, doc *colly.XMLElement) (item *video.VideoInfo) {
	var (
		ele   = fetcher.NewXMLElement(doc)
		cover = "src"
	)
	if source.SourceKey == "beqegecc" {
		cover = "data-original"
	}
	item = &video.VideoInfo{
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

func (s *fetcherService) parseItemSearch(source *video.VideoSource, doc *colly.XMLElement) (item *video.VideoInfo) {
	var ele = fetcher.NewXMLElement(doc)
	item = &video.VideoInfo{
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

func (s *fetcherService) parseItemInfo(source *video.VideoSource, doc *colly.XMLElement) (info video.VideoInfo) {
	var ele = fetcher.NewXMLElement(doc)
	info = video.VideoInfo{
		Name:         ele.ChildText(source.DetailVideoNameRule),
		Author:       ele.ChildText(source.DetailVideoAuthorRule),
		Cover:        ele.ChildUrlText(source.DetailVideoCoverRule),
		Category:     ele.ChildText(source.DetailVideoCategoryRule),
		Description:  ele.ChildHtml(source.DetailVideoDescriptionRule),
		NewChapter:   ele.ChildText(source.DetailVideoNewChapterRule),
		URL:          ele.ChildUrl(source.DetailVideoNewChapterUrlRule, "href"),
		FirstChapter: ele.ChildText(source.DetailVideoFirstChapterRule),
		FirstURL:     ele.ChildUrl(source.DetailVideoFirstUrlRule, "href"),
		Source:       source.SourceKey,
	}
	for _, v := range auts {
		if strings.Contains(info.Author, v) {
			info.Author = strings.Split(info.Author, v)[1]
		}
	}
	return
}

func (s *fetcherService) parseChapterList(source *video.VideoSource, doc *colly.XMLElement, url string) (chapter video.Chapter) {
	var ele = fetcher.NewXMLElement(doc)
	chapter = video.Chapter{
		Title:      ele.ChildText(source.DetailChapterTitleRule),
		DetailURL:  url,
		ChapterURL: ele.ChildUrl(source.DetailChapterURLRule, "href"),
		Source:     source.SourceKey,
	}

	return
}

func (s *fetcherService) parseContent(source *video.VideoSource, doc *colly.XMLElement, url string) (content video.VideoContent) {
	var ele = fetcher.NewXMLElement(doc)
	content = video.VideoContent{
		Title:       ele.ChildText(source.ContentTitleRule),
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
		content.Text = ele.ChildRemoveHtml(source.ContentTextRule, "div")
	//case "biqugee":
	//	content.Text = ele.ChildHtml(source.ContentTextRule)
	default:
		content.Text = ele.ChildHtml(source.ContentTextRule)
		i := strings.Index(content.Text, "<p><a")
		k := strings.Index(content.Text, "a></p>")
		if i != -1 {
			content.Text = content.Text[:i] + content.Text[k+6:]
		}
	}
	return content
}
