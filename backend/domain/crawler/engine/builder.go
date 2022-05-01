package engine

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/fetcher"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/parser"
)

type CrawlerBuilder interface {
	HatenaHotEntry() *CrawlerEngine
	QiitaEntry() *CrawlerEngine
}

func NewCrawlerBuilder() CrawlerBuilder {
	return builderImpl{}
}

type builderImpl struct{}

func (b builderImpl) HatenaHotEntry() *CrawlerEngine {
	fetcher := fetcher.New(fetcher.HatenaHotEntryLogic{})
	parsre := parser.New(parser.HatenaHotEntryLogic{})
	return NewCrawlerEngine("HatenaHotEntryCrawler", fetcher, parsre)
}

func (b builderImpl) QiitaEntry() *CrawlerEngine {
	fetcher := fetcher.New(fetcher.QiitaEntryLogic{})
	parsre := parser.New(parser.QiitaEntryLogic{})
	return NewCrawlerEngine("QiitaEntryCrawler", fetcher, parsre)
}
