package datasource

import (
	"dnovel/models/datamodels"
	"dnovel/util/cons"
)

// XPATH: 语法
var BookSources = map[int64]datamodels.BookSource{
	1: {
		SourceName: "笔趣阁",
		SourceKey:  "biquha",

		ClassifyUrl: map[string][]string{
			//cons.IntelRec:  {"https://www.sobiquge.com/"},
			cons.Classify1: {"http://m.biquha.com/xuanhuan/"},
			cons.Classify2: {"http://m.biquha.com/wuxia/"},
			cons.Classify3: {"http://m.biquha.com/kehuan/"},
			cons.Classify4: {"http://m.biquha.com/wangyou/"},
			cons.Classify5: {"http://m.biquha.com/lishi/"},
			cons.Classify6: {"http://m.biquha.com/dushi/"},
		},
		ClassifyItemRule:   `//div[@class="hot"]/div`,
		ClassifyItemName:   `//div[@class="item"]/div[1]/dl/dt/a`,
		ClassifyItemAuthor: `//div[@class="item"]/div[1]/dl/dt/span`,
		ClassifyItemUrl:    `//div[@class="item"]/div[1]/dl/dt/a`,
		ClassifyItemCover:  `//div[@class="item"]/div[1]/div[1]/a/img`,
		ClassifyItemDesc:   `//div[@class="item"]/div[1]/dl/dd`,

		SearchURL:                "http://m.biquha.com/drhugirdhugi.php?q=%s",
		SearchItemRule:           `//div[@class="wrap"]/div[1]/div`,
		SearchItemNameRule:       `./div[2]/h4/a`,
		SearchItemAuthorRule:     `./div[2]/div[2]`,
		SearchItemCoverRule:      `./div[1]/a/img`,
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: `./div[2]/div[3]/a`,
		SearchItemURLRule:        `./div[2]/h4/a`,

		DetailBookItemRule:          `//body`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[3]`,
		DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
		DetailBookCategoryRule:      `//div[@id="info"]/p[1]`,
		DetailBookDescriptionRule:   `//div[@id="intro"]/p[1]`,
		DetailBookNewChapterRule:    `//div[@id="info"]/p[5]/a`,
		DetailBookNewChapterUrlRule: `//div[@id="info"]/p[5]/a`,
		DetailBookFirstChapterRule:  `//div[@class="listmain"]/dl/dd[1]/a`,
		DetailBookFirstUrlRule:      `//div[@class="listmain"]/dl/dd[1]/a`,

		DetailChapterRule:      `//div[@class="listmain"]/dl/dd`, //  `//div[@class="listmain"]/dl/dd[position()>=13]`
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule:   `./a`,

		ContentTitleRule:       `//div[@class="content"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="content"]/div[1]/ul/li[1]/a`,
		ContentNextURLRule:     `//div[@class="content"]/div[1]/ul/li[3]/a`,
	},
	2: {
		SourceName: "新笔趣阁",
		SourceKey:  "xbiquge", // 等同www.ibiquge.la

		ClassifyUrl: map[string][]string{
			cons.Classify1: {"https://www.ibiquges.org//xuanhuanxiaoshuo/"},
			cons.Classify2: {"https://www.ibiquges.org//xiuzhenxiaoshuo/"},
			cons.Classify3: {"https://www.ibiquges.org//kehuanxiaoshuo/"},
			cons.Classify4: {"https://www.ibiquges.org//wangyouxiaoshuo/"},
			cons.Classify5: {"https://www.ibiquges.org//chuanyuexiaoshuo/"},
			cons.Classify6: {"https://www.ibiquges.org//dushixiaoshuo/"},
		},
		ClassifyItemRule:   `//div[@id="newscontent"]/div[1]/ul/li`,
		ClassifyItemName:   `//span/a`,
		ClassifyItemAuthor: `//span[3]`,
		ClassifyItemUrl:    `//span/a`,
		//ClassifyItemCover:  `//div[@id="fmimg"]/img`,
		//ClassifyItemDesc:   `//div[@id="intro"]/p[2]`,

		SearchURL:                "https://www.ibiquges.org//modules/article/waps.php?searchkey=%s",
		SearchItemRule:           "//tbody/tr[1]/following-sibling::tr",
		SearchItemNameRule:       "./td[1]/a",
		SearchItemAuthorRule:     "./td[3]",
		SearchItemCoverRule:      "",
		SearchItemCategoryRule:   "",
		SearchItemNewChapterRule: "./td[2]/a",
		SearchItemURLRule:        "./td[1]/a",

		DetailBookItemRule:          `//div[@id="wrapper"]`,
		DetailBookNameRule:          `//div[@id="info"]/h1`,
		DetailBookAuthorRule:        `//div[@id="info"]/p[1]`,
		DetailBookCoverRule:         `//div[@id="fmimg"]/img/@src`,
		DetailBookCategoryRule:      `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule:   `//div[@id="intro"]/p[2]`,
		DetailBookNewChapterRule:    `//div[@id="info"]/p[4]/a`,
		DetailBookNewChapterUrlRule: `//div[@id="info"]/p[4]/a`,
		DetailBookFirstChapterRule:  `//div[@id="list"]/dl/dd[1]/a`,
		DetailBookFirstUrlRule:      `//div[@id="list"]/dl/dd[13]/a`,

		DetailChapterRule:      `//div[@id="list"]/dl/dd`,
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule:   `./a`,

		ContentTitleRule:       `//div[@class="bookname"]/h1`,
		ContentTextRule:        `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule:     `//div[@class="bottem1"]/a[4]`,
	},
	3: {
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
		ClassifyItemCover:  `//a[@class="bookimg"]/img`, // 403
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
		DetailBookCoverRule:         `//div[@id="sidebar"]/div[@id="fmimg"]/img/@data-original]`, // 403
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
}
