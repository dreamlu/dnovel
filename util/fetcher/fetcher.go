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
			r.Headers.Set("cookie", "cf_clearance=8f7e8826a9fca60e000b254aa1c62eb66ea95e0d-1667200369-0-150; __cf_bm=5Cm7Fkpqq9.1MnLU2BazDmowXaqnhikXmo01Sdt3zvM-1667205564-0-AdSmEFZDmispSpenxu3ozQhp2WaM7fT6qNTGeiPR4/R/h14+fc/6d0sHdrnuedOcTs5ADq8P3vlOefRTbw6Nq2WyCsamhMZr0qaWIzObr0ieRp1qc1V0yFA03zvwfBb9aQ==")
		}
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
