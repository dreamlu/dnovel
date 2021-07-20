package datasource

import (
	"dnovel/models/datamodels"
	"dnovel/util/cons"
)

// XPATH: 语法
var BookSources = map[int64]datamodels.BookSource{
	1: {
		SourceName: "笔趣阁cc",
		SourceKey:  "beqegecc",

		ClassifyUrl: map[string][]string{
			cons.IntelRec:  {"http://www.beqege.cc"},
			cons.Classify1: {"http://www.beqege.cc/class1/"},
			cons.Classify2: {"http://www.beqege.cc/class2/"},
			cons.Classify3: {"http://www.beqege.cc/class6/"},
			cons.Classify4: {"http://www.beqege.cc/class5/"},
			cons.Classify5: {"http://www.beqege.cc/class4/"},
			cons.Classify6: {"http://www.beqege.cc/class3/"},
		},
		ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
		ClassifyItemName:   `//span/a`,
		ClassifyItemAuthor: `//span[@class="s4"] or span[@class="s4"]`,
		ClassifyItemUrl:    `//span/a`,

		SearchURL:            "http://www.beqege.cc/search.php?keyword=%s",
		SearchItemRule:       `//ul[@class="chapter-list"]/li`,
		SearchItemNameRule:   `./span[@class="s2"]/a`,
		SearchItemAuthorRule: `./span[@class="s4"]`,
		//SearchItemCoverRule:      "", // 不需要
		SearchItemCategoryRule:   `./span[@class="s1"]`,
		SearchItemNewChapterRule: `./span[@class="s3"]/a`,
		SearchItemURLRule:        `./span[@class="s2"]/a`,

		DetailBookItemRule:          `//div[@id="main"]`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:         `//div[@id="sidebar"]/div[@id="fmimg"]/img/@data-original]`,
		DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule:   `//div[@id="intro"]/p[1]`,
		DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
		DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
		DetailBookFirstChapterRule:  `//div[@id="list"]/dl/dd[1]/a`,
		DetailBookFirstUrlRule:      `//div[@id="list"]/dl/dd[1]/a`,

		DetailChapterRule:      `//div[@id="list"]/dl/dd`,
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule:   `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[1]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[3]`,
	},
	//2: {
	//	SourceName: "笔趣阁",
	//	SourceKey:  "biquwx",
	//
	//	ClassifyUrl: map[string][]string{
	//		cons.IntelRec:  {"https://www.biquwx.la/"},
	//		cons.Classify1: {"https://www.biquwx.la/list/1_1.html"},
	//		cons.Classify2: {"https://www.biquwx.la/list/2_1.html"},
	//		cons.Classify3: {"https://www.biquwx.la/list/6_1.html"},
	//		cons.Classify4: {"https://www.biquwx.la/list/5_1.html"},
	//		cons.Classify5: {"https://www.biquwx.la/list/4_1.html"},
	//		cons.Classify6: {"https://www.biquwx.la/list/3_1.html"},
	//	},
	//	ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
	//	ClassifyItemName:   `//span/a`,
	//	ClassifyItemAuthor: `//span[4]`,
	//	ClassifyItemUrl:    `//span/a`,
	//
	//	SearchURL:                "https://www.biquwx.la/modules/article/search.php?searchkey=%s",
	//	SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
	//	SearchItemNameRule:       "./td[1]/a",
	//	SearchItemAuthorRule:     "./td[3]",
	//	SearchItemCoverRule:      "",
	//	SearchItemCategoryRule:   "",
	//	SearchItemNewChapterRule: "./td[2]/a",
	//	SearchItemURLRule:        "//td[1]/a",
	//
	//	DetailBookItemRule:          `//div[@id="wrapper"]`,
	//	DetailBookNameRule:          `//div[@id="info"]/h1`,
	//	DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
	//	DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
	//	DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
	//	DetailBookDescriptionRule:   `//div[@id="intro"]/p[1]`,
	//	DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
	//	DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
	//	DetailBookFirstChapterRule:  `//div[@id="list"]/dl/dd[1]/a`,
	//	DetailBookFirstUrlRule:      `//div[@id="list"]/dl/dd[1]/a`,
	//
	//	DetailChapterRule:      `//div[@id="list"]/dl/dd`,
	//	DetailChapterTitleRule: `./a`,
	//	DetailChapterURLRule:   `./a`,
	//
	//	ContentTitleRule:       `//div[@class="bookname"]/h1`,
	//	ContentTextRule:        `//div[@id="content"]`,
	//	ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
	//	ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	//},
	2: {
		SourceName: "笔趣阁",
		SourceKey:  "sobiquge",

		ClassifyUrl: map[string][]string{
			cons.IntelRec:  {"https://www.sobiquge.com/"},
			cons.Classify1: {"https://www.sobiquge.com/xuanhuan"},
			cons.Classify2: {"https://www.sobiquge.com/xiuzhen"},
			cons.Classify3: {"https://www.sobiquge.com/kehuan"},
			cons.Classify4: {"https://www.sobiquge.com/wangyou"},
			cons.Classify5: {"https://www.sobiquge.com/lishi"},
			cons.Classify6: {"https://www.sobiquge.com/dushi"},
		},
		ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
		ClassifyItemName:   `//span/a`,
		ClassifyItemAuthor: `//span[@class="s4"] or span[@class="s4"]`,
		ClassifyItemUrl:    `//span/a`,

		SearchURL:                "https://m.sobiquge.com/search.php?q=%s",
		SearchItemRule:           `//div[@class="result-item result-game-item"]`,
		SearchItemNameRule:       "./div[2]/h3/a/span",
		SearchItemAuthorRule:     "./div[2]/div/p[1]/span[2]",
		SearchItemCoverRule:      `./div[1]/a/img`,
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./div[2]/div/p[4]/a",
		SearchItemURLRule:        "./div[2]/h3/a",

		DetailBookItemRule:          `//div[@id="wrapper"]`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
		DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule:   `//div[@id="intro"]`,
		DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
		DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
		DetailBookFirstChapterRule:  `//div[@id="list"]/dl/dd[1]/a`,
		DetailBookFirstUrlRule:      `//div[@id="list"]/dl/dd[1]/a`,

		DetailChapterRule:      `//div[@id="list"]/dl/dd`,
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule:   `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[1]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[3]`,
	},
}
