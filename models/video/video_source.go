package video

type VideoSource struct {
	SourceName string `json:"source_name" comment:"书源名称"`
	SourceKey  string `json:"source_key" comment:"书源标识"`

	// 分类模块规则
	ClassifyUrl        map[string][]string `json:"classify_url"`         // 分类url列表
	ClassifyItemRule   string              `json:"classify_item_rule"`   // 分类查找规则
	ClassifyItemName   string              `json:"classify_item_name"`   // 动漫名称
	ClassifyItemAuthor string              `json:"classify_item_author"` // 作者
	ClassifyItemUrl    string              `json:"classify_item_url"`    // 书url
	ClassifyItemCover  string              `json:"classify_item_cover"`  // 书封面
	ClassifyItemDesc   string              `json:"classify_item_desc"`   // 书籍描述

	SearchURL                string `json:"search_url" comment:"搜索网址"`
	SearchItemRule           string `json:"search_list_rule" comment:"搜索结果规则"`
	SearchItemNameRule       string `json:"search_item_name_rule" comment:"搜索结果名称规则"`
	SearchItemAuthorRule     string `json:"search_item_author_rule" comment:"搜索结果作者规则"`
	SearchItemCoverRule      string `json:"search_item_cover_rule" comment:"搜索结果封面规则"`
	SearchItemCategoryRule   string `json:"search_item_category_rule" comment:"搜索结果分类规则"`
	SearchItemNewChapterRule string `json:"search_item_new_chapter_rule" comment:"搜索结果最新章节规则"`
	SearchItemURLRule        string `json:"search_item_url_rule" comment:"搜索结果链接规则"`

	DetailVideoItemRule          string `json:"detail_book_item_rule" comment:"动漫详情规则"`
	DetailVideoNameRule          string `json:"detail_book_name_rule" comment:"动漫名称规则"`
	DetailVideoAuthorRule        string `json:"detail_book_author_rule" comment:"动漫作者规则"`
	DetailVideoCoverRule         string `json:"detail_book_cover_rule" comment:"动漫封面规则"`
	DetailVideoCategoryRule      string `json:"detail_book_category_rule" comment:"动漫分类规则"`
	DetailVideoDescriptionRule   string `json:"detail_book_description_rule" comment:"动漫描述规则"`
	DetailVideoNewChapterRule    string `json:"detail_book_new_chapter_rule" comment:"动漫最新章节规则"`
	DetailVideoNewChapterUrlRule string `json:"detail_book_new_chapter_url_rule" comment:"动漫最新章节url规则"`
	DetailVideoFirstChapterRule  string `json:"detail_book_first_chapter_rule"` // 第一章规则
	DetailVideoFirstUrlRule      string `json:"detail_book_first_url_rule"`     // 第一章url规则

	DetailChapterRule      string `json:"detail_chapter_list_rule" comment:"动漫章节列表规则"`
	DetailChapterTitleRule string `json:"detail_chapter_title_rule" comment:"动漫章节名称规则"`
	DetailChapterURLRule   string `json:"detail_chapter_url_rule" comment:"动漫章节链接规则"`

	ContentTitleRule       string `json:"chapter_content_rule" comment:"内容标题规则"`
	ContentTextRule        string `json:"content_text_rule" comment:"内容正文规则"`
	ContentPreviousURLRule string `json:"chapter_previous_url_rule" comment:"内容上一章链接规则"`
	ContentNextURLRule     string `json:"chapter_next_url_rule" comment:"内容下一章链接规则"`

	//Weight int `json:"weight" comment:"权重"`
}
