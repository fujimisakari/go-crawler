package service

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlaggregate/entity"
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
)

func GetCrawlAggregate(crawlDate string) (*entity.CrawlAggregate, error) {
	crawlEntryRepo := crawlentry.NewRepository()
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		return nil, err
	}

	hatenaHotEntryRepo := hatenahotentry.NewRepository()
	hatenaHotEntryList, err := hatenaHotEntryRepo.FindByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		return nil, err
	}

	qiitaEntryRepo := qiitaentry.NewRepository()
	qittaEntryList, err := qiitaEntryRepo.FindByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"date":               crawlDate,
		"hatenaHotEntryList": hatenaHotEntryList,
		"qiitaEntryList":     qittaEntryList,
	}
	return entity.NewCrawlAggregate(params), nil
}
