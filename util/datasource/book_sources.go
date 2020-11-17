package datasource

import "dnovel/models/datamodels"

var BookSources = map[int64]datamodels.BookSource{
	1: {
		SourceName: "酷小说",
		SourceURL:  "https://www.kuxiaoshuo.com",
		SourceKey:  "kuxiaoshuo",

		SearchURL:                "https://www.kuxiaoshuo.com/modules/article/search.php?searchkey=%s",
		SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
		SearchItemNameRule:       "./td[1]/a",
		SearchItemAuthorRule:     "./td[3]",
		SearchItemCoverRule:      "",
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./td[2]/a",
		SearchItemURLRule:        "./td[1]/a",

		DetailBookItemRule:        `//div[@id="wrapper"]`,
		DetailBookNameRule:        `//div[@id="info"]/h1`,
		DetailBookAuthorRule:      `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:       `//div[@id="fmimg"]/img`,
		DetailBookCategoryRule:    `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule: `//div[@id="intro"]/p[1]`,
		DetailBookNewChapterRule:  "",

		DetailChapterListURLRule:  "",
		DetailNewChapterRule:      `//table[@id="adt2"]/parent::div/preceding-sibling::dd`,
		DetailNewChapterTitleRule: `./a`,
		DetailNewChapterURLRule:   `./a`,
		DetailChapterRule:         `//dt[2]/following-sibling::dd`,
		DetailChapterTitleRule:    `./a`,
		DetailChapterURLRule:      `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
	2: {
		SourceName: "笔趣阁",
		SourceURL:  "http://www.biquge.info",
		SourceKey:  "biquge",

		SearchURL:                "http://www.biquge.info/modules/article/search.php?searchkey=%s",
		SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
		SearchItemNameRule:       "./td[1]/a",
		SearchItemAuthorRule:     "./td[3]",
		SearchItemCoverRule:      "",
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./td[2]/a",
		SearchItemURLRule:        "//td[1]/a",

		DetailBookItemRule:        `//div[@id="wrapper"]`,
		DetailBookNameRule:        `//div[@id="info"]/h1`,
		DetailBookAuthorRule:      `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:       `//div[@id="fmimg"]/img`,
		DetailBookCategoryRule:    `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule: `//div[@id="intro"]/p[1]`,
		DetailBookNewChapterRule:  "",

		DetailChapterListURLRule:  "",
		DetailNewChapterRule:      "",
		DetailNewChapterTitleRule: "",
		DetailNewChapterURLRule:   "",
		DetailChapterRule:         `//div[@id="list"]/dl/dd`,
		DetailChapterTitleRule:    `./a`,
		DetailChapterURLRule:      `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
	3: {
		SourceName:               "笔趣阁【vipzw】",
		SourceURL:                "http://www.vipzw.com",
		SourceKey:                "vipzw",
		SearchURL:                "http://www.vipzw.com/search.php?searchkey=%s",
		SearchItemRule:           `//div[@class="item"]`,
		SearchItemNameRule:       "//dt/a",
		SearchItemAuthorRule:     "//dt/span",
		SearchItemCoverRule:      "//img",
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "",
		SearchItemURLRule:        `//dt/a`,

		DetailBookItemRule:        `//div[@id="wrapper"]`,
		DetailBookNameRule:        `//div[@id="info"]/h1`,
		DetailBookAuthorRule:      `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:       `//div[@id="fmimg"]/img`,
		DetailBookCategoryRule:    `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule: `//div[@id="intro"]/p[1]`,
		DetailChapterListURLRule:  "",
		DetailNewChapterRule:      `//table[@id="adt2"]/parent::div/preceding-sibling::dd`,
		DetailNewChapterTitleRule: `./a`,
		DetailNewChapterURLRule:   `./a`,
		DetailChapterRule:         `//dt[2]/following-sibling::dd`,
		DetailChapterTitleRule:    `./a`,
		DetailChapterURLRule:      `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
}
