package fetcher

import (
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/util/cons"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/gocolly/colly/v2/extensions"
	"log"
)

// 提取器
func NewFetcher() *colly.Collector {
	var c *colly.Collector
	if conf.GetString("app.devMode") != cons.Prod {
		c = colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	} else {
		c = colly.NewCollector()
	}

	extensions.Referer(c)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
