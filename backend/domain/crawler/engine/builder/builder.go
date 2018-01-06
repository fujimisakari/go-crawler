package builder

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/fetcher"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/parser"
)

type (
	HatenaHotEntryBuilder struct{}
	QiitaEntryBuilder     struct{}
)

func (b HatenaHotEntryBuilder) BuildCrawlerEngine() *engine.CrawlerEngine {
	fetcher := fetcher.New(fetcher.HatenaHotEntryLogic{})
	parsre := parser.New(parser.HatenaHotEntryLogic{})
	return engine.New("HatenaHotEntryCrawler", fetcher, parsre)
}

func (b QiitaEntryBuilder) BuildCrawlerEngine() *engine.CrawlerEngine {
	fetcher := fetcher.New(fetcher.QiitaEntryLogic{})
	parsre := parser.New(parser.QiitaEntryLogic{})
	return engine.New("QiitaEntryCrawler", fetcher, parsre)
}
