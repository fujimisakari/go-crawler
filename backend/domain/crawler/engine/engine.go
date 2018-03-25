package engine

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/fetcher"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/parser"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/result"
)

type CrawlerEngine struct {
	Name        string
	fetchEngine *fetcher.FetchEngine
	parseEngine *parser.ParseEngine
}

func NewCrawlerEngine(name string, fetchEngine *fetcher.FetchEngine, parseEngine *parser.ParseEngine) *CrawlerEngine {
	return &CrawlerEngine{
		Name:        name,
		fetchEngine: fetchEngine,
		parseEngine: parseEngine,
	}
}

func (c *CrawlerEngine) Run(cDone chan<- *result.CrawlResult) {
	rDone := response.NewChanel()
	defer close(rDone)

	go c.fetchEngine.Fetch(rDone)
	res := <-rDone
	if res.Err != nil {
		cDone <- result.New(nil, res.Err)
		return
	}
	entries, err := c.parseEngine.Parse(res)
	if err != nil {
		cDone <- result.New(nil, res.Err)
		return
	}
	cDone <- result.New(entries, nil)
}
