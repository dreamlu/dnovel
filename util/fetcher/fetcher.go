package fetcher

import (
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/util/cons"
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
	if conf.GetString("app.devMode") != cons.Prod {
		c = colly.NewCollector(dg)
	} else {
		c = colly.NewCollector()
	}

	extensions.Referer(c)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		if strings.Contains(r.URL.String(), "qxzx8") {
			r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
			r.Headers.Set("accept-encoding", "gzip, deflate, br")
			r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
			r.Headers.Set("cache-control", "no-cache")
			r.Headers.Set("pragma", "no-cache")
			r.Headers.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
			r.Headers.Set("sec-ch-ua-mobile", "?0")
			r.Headers.Set("sec-ch-ua-platform", "Linux")
			r.Headers.Set("sec-fetch-dest", "document")
			r.Headers.Set("sec-fetch-mode", "navigate")
			r.Headers.Set("sec-fetch-site", "none")
			r.Headers.Set("sec-fetch-user", "?1")
			r.Headers.Set("upgrade-insecure-requests", "1")
			r.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
		}
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
