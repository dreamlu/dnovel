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
	dg := colly.Debugger(&debug.LogDebugger{})
	ag := colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	if conf.GetString("app.devMode") != cons.Prod {
		c = colly.NewCollector(dg, ag)
	} else {
		c = colly.NewCollector(ag)
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
