package datasource

import (
	"dnovel/models/datamodels"
	"dnovel/util/cons"
)

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
		ClassifyItemAuthor: `//span[3]`,
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
	2: {
		SourceName: "笔趣阁",
		SourceKey:  "biquge",

		ClassifyUrl: map[string][]string{
			cons.IntelRec:  {"http://www.biquge.info"},
			cons.Classify1: {"http://www.biquge.info/list/1_1.html"},
			cons.Classify2: {"http://www.biquge.info/list/2_1.html"},
			cons.Classify3: {"http://www.biquge.info/list/6_1.html"},
			cons.Classify4: {"http://www.biquge.info/list/5_1.html"},
			cons.Classify5: {"http://www.biquge.info/list/4_1.html"},
			cons.Classify6: {"http://www.biquge.info/list/3_1.html"},
		},
		ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
		ClassifyItemName:   `//span/a`,
		ClassifyItemAuthor: `//span[4]`,
		ClassifyItemUrl:    `//span/a`,

		SearchURL:                "http://www.biquge.info/modules/article/search.php?searchkey=%s",
		SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
		SearchItemNameRule:       "./td[1]/a",
		SearchItemAuthorRule:     "./td[3]",
		SearchItemCoverRule:      "",
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./td[2]/a",
		SearchItemURLRule:        "//td[1]/a",

		DetailBookItemRule:          `//div[@id="wrapper"]`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
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
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
	//3: {
	//	SourceName: "新笔趣阁",
	//	SourceKey:  "xbiquge",
	//
	//	ClassifyUrl: map[string][]string{
	//		cons.Classify1: {"http://www.xbiquge.la/xuanhuanxiaoshuo/"},
	//		cons.Classify2: {"http://www.xbiquge.la/xiuzhenxiaoshuo/"},
	//		cons.Classify3: {"http://www.xbiquge.la/kehuanxiaoshuo/"},
	//		cons.Classify4: {"http://www.xbiquge.la/wangyouxiaoshuo/"},
	//		cons.Classify5: {"http://www.xbiquge.la/chuanyuexiaoshuo/"},
	//		cons.Classify6: {"http://www.xbiquge.la/dushixiaoshuo/"},
	//	},
	//	ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
	//	ClassifyItemName:   `//span/a`,
	//	ClassifyItemAuthor: `//span[3]`,
	//	ClassifyItemUrl:    `//span/a`,
	//
	//	SearchURL:                "http://www.xbiquge.la/modules/article/waps.php?searchkey=%s",
	//	SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
	//	SearchItemNameRule:       "./td[1]/a",
	//	SearchItemAuthorRule:     "./td[3]",
	//	SearchItemCoverRule:      "",
	//	SearchItemCategoryRule:   "",
	//	SearchItemNewChapterRule: "./td[2]/a",
	//	SearchItemURLRule:        "./td[1]/a",
	//
	//	DetailBookItemRule:          `//div[@id="wrapper"]`,
	//	DetailBookNameRule:          `//div[@id="info"]/h1`,
	//	DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
	//	DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
	//	DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
	//	DetailBookDescriptionRule:   `//div[@id="intro"]/p[2]`,
	//	DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
	//	DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
	//
	//	DetailChapterRule:      `//dt[2]/following-sibling::dd`,
	//	DetailChapterTitleRule: `./a`,
	//	DetailChapterURLRule:   `./a`,
	//
	//	ContentTitleRule:       `//div[@class="bookname"]/h1`,
	//	ContentTextRule:        `//div[@id="content"]`,
	//	ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
	//	ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	//},
}
