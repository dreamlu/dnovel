package fetcher

import (
	"github.com/dreamlu/gt/conf"
	"github.com/dreamlu/gt/crud/dep/cons"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/gocolly/colly/v2/extensions"
	"log"
	"strings"
)

// 提取器
func NewFetcher() *colly.Collector {
	var c *colly.Collector
	dg := colly.Debugger(&debug.LogDebugger{})
	//ag := colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	if conf.Get[string]("app.devMode") != cons.Prod {
		c = colly.NewCollector(dg)
	} else {
		c = colly.NewCollector()
	}

	extensions.Referer(c)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		if strings.Contains(r.URL.String(), "beqegecc") {
			r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
			r.Headers.Set("cookie", "cf_clearance=6f520d523c83408ee2ccbb81fab2588ac9b8352a-1667284728-0-150")
		}
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
