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
			//cons.IntelRec:  {"http://m.beqege.cc"},
			cons.Classify1: {"https://m.beqege.cc/class1/"},
			cons.Classify2: {"https://m.beqege.cc/class2/"},
			cons.Classify3: {"https://m.beqege.cc/class6/"},
			cons.Classify4: {"https://m.beqege.cc/class5/"},
			cons.Classify5: {"https://m.beqege.cc/class4/"},
			cons.Classify6: {"https://m.beqege.cc/class3/"},
		},
		ClassifyItemRule:   `//div[@class="recommend"]/div[@class="booklist border-bottom"]`,
		ClassifyItemName:   `//div[@class="bookinfo"]/div[1]/a`,
		ClassifyItemAuthor: `//div[@class="bookinfo"]/div[2]`, //`//span[@class="s4"] or span[@class="s4"]`,
		ClassifyItemUrl:    `//div[@class="bookinfo"]/div[1]/a`,
		ClassifyItemCover:  `//a[@class="bookimg"]/img`,
		ClassifyItemDesc:   `//div[@class="bookinfo"]/div[3]`,

		SearchURL:            "https://www.beqege.cc/search.php?keyword=%s",
		SearchItemRule:       `//ul[@class="chapter-list"]/li`,
		SearchItemNameRule:   `./span[@class="s2"]/a`,
		SearchItemAuthorRule: `./span[@class="s4"]`,
		//SearchItemCoverRule:      "", // 搜索结果没有封面
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
		SourceKey:  "ibiquge",

		ClassifyUrl: map[string][]string{
			//cons.IntelRec:  {"https://www.sobiquge.com/"},
			cons.Classify1: {"https://m.ibiquge.net/xclass/1/1.html"},
			cons.Classify2: {"https://m.ibiquge.net/xclass/2/1.html"},
			cons.Classify3: {"https://m.ibiquge.net/xclass/6/1.html"},
			cons.Classify4: {"https://m.ibiquge.net/xclass/5/1.html"},
			cons.Classify5: {"https://m.ibiquge.net/xclass/4/1.html"},
			cons.Classify6: {"https://m.ibiquge.net/xclass/3/1.html"},
		},
		ClassifyItemRule:   `//div[@id="main"]/div`,
		ClassifyItemName:   `//div[@class="bookinfo"]/div[1]/a/p`,
		ClassifyItemAuthor: `//div[@class="bookinfo"]/div[1]/p`,
		ClassifyItemUrl:    `//div[@class="bookimg"]/a[1]`,
		ClassifyItemCover:  `//div[@class="bookimg"]/a[1]/img`,
		ClassifyItemDesc:   `//p[@class="review"]`,

		SearchURL:                "https://m.ibiquge.net/SearchBook.php?keyword=%s",
		SearchItemRule:           `//div[@class="slide-item list1"]/div[@class="hot_sale"]`,
		SearchItemNameRule:       "./a[1]/p",
		SearchItemAuthorRule:     "./p[1]",
		SearchItemCoverRule:      `./p`,
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./p[2]/a",
		SearchItemURLRule:        "./a[1]",

		DetailBookItemRule:          `//div[@id="main"]`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
		DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule:   `//div[@id="intro"]`,
		DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
		DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
		DetailBookFirstChapterRule:  `//div[@id="list"]/dl/dd[1]/a`,
		DetailBookFirstUrlRule:      `//div[@id="list"]/dl/dd[13]/a`,

		DetailChapterRule:      `//div[@id="list"]/dl/dd[position()>=13]`, // position()>13 and position()<15
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule:   `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
	// todo remove part use
	//3: {
	//	SourceName: "笔趣阁",
	//	SourceKey:  "biqugee",
	//
	//	ClassifyUrl: map[string][]string{
	//		//cons.IntelRec:  {"https://www.sobiquge.com/"},
	//		cons.Classify1: {"https://m.biqugee.com/xuanhuan"},
	//		cons.Classify2: {"https://m.biqugee.com/xiuzhen"},
	//		cons.Classify3: {"https://m.biqugee.com/kehuan"},
	//		cons.Classify4: {"https://m.biqugee.com/wangyou"},
	//		cons.Classify5: {"https://m.biqugee.com/lishi"},
	//		cons.Classify6: {"https://m.biqugee.com/dushi"},
	//	},
	//	ClassifyItemRule:   `//div[@class="bookbox"]`,
	//	ClassifyItemName:   `//div[@class="bookinfo"]/h4/i/a`,
	//	ClassifyItemAuthor: `//div[@class="bookinfo"]/div[1]`,
	//	ClassifyItemUrl:    `//div[@class="bookinfo"]/h4/i/a`,
	//	ClassifyItemCover:  `//div[@class="bookimg"]/a/img`,
	//	ClassifyItemDesc:   `//div[@class="intro_line"]`,
	//
	//	SearchURL:                "https://m.biqugee.com/search.php?q=%s",
	//	SearchItemRule:           `//div[@class="result-item result-game-item"]`,
	//	SearchItemNameRule:       "./div[2]/h3/a/span",
	//	SearchItemAuthorRule:     "./div[2]/div/p[1]/span[2]",
	//	SearchItemCoverRule:      `./div[1]/a/img`,
	//	SearchItemCategoryRule:   "",
	//	SearchItemNewChapterRule: "./div[2]/div/p[4]/a",
	//	SearchItemURLRule:        "./div[2]/h3/a",
	//
	//	DetailBookItemRule:          `//div[@id="wrapper"]`,
	//	DetailBookNameRule:          `//div[@id="info"]/h1`,
	//	DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
	//	DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
	//	DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
	//	DetailBookDescriptionRule:   `//div[@id="intro"]`,
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
	//	ContentPreviousURLRule: `//div[@class="bottem1"]/a[1]`,
	//	ContentNextURLRule:     `//div[@class="bottem1"]/a[3]`,
	//},
}
