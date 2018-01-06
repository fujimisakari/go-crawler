package builder

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/fetcher"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/parser"
)

type Builder interface {
	HatenaHotEntry() *engine.CrawlerEngine
	QiitaEntry() *engine.CrawlerEngine
}

func New() Builder {
	return builderImpl{}
}

type builderImpl struct{}


func (b builderImpl) HatenaHotEntry() *engine.CrawlerEngine {
	fetcher := fetcher.New(fetcher.HatenaHotEntryLogic{})
	parsre := parser.New(parser.HatenaHotEntryLogic{})
	return engine.New("HatenaHotEntryCrawler", fetcher, parsre)
}

func (b builderImpl) QiitaEntry() *engine.CrawlerEngine {
	fetcher := fetcher.New(fetcher.QiitaEntryLogic{})
	parsre := parser.New(parser.QiitaEntryLogic{})
	return engine.New("QiitaEntryCrawler", fetcher, parsre)
}
