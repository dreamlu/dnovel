package video

type VideoInfo struct {
	Name         string `json:"name" bson:"name" comment:"动漫名称"`
	Author       string `json:"author" bson:"author" comment:"动漫作者"`
	Cover        string `json:"cover" bson:"cover" comment:"动漫封面"`
	Category     string `json:"category" bson:"category" comment:"动漫分类"`
	Description  string `json:"description" bson:"description" comment:"动漫描述"`
	NewChapter   string `json:"new_chapter" bson:"new_chapter" comment:"搜索结果最新章节"`
	URL          string `json:"url" bson:"url" comment:"搜索结果链接"`
	FirstChapter string `json:"first_chapter"` // 第一章
	FirstURL     string `json:"first_url"`     // 第一章url
	Source       string `json:"source" bson:"source" comment:"搜索结果来源"`
}
