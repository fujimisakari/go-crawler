package engine

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/fetcher"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/parser"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
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

func (c *CrawlerEngine) Run() ([]map[string]interface{}, error) {
	content, err := c.fetchEngine.Fetch()
	if err != nil {
		return nil, err
	}
	response := response.New(content)
	entries, err := c.parseEngine.Parse(response)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
