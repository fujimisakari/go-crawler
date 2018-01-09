package qiitaentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

func GetByID(id int) (*entity.QiitaEntry, error) {
	qiitaEntryRepo := qiitaentry.NewRepository()
	qittaEntry, err := qiitaEntryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return qittaEntry, nil
}

func GetListByDate(crawlDate string) (*entity.QiitaEntryList, error) {
	crawlEntryRepo := crawlentry.NewRepository()
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		return nil, err
	}

	qiitaEntryRepo := qiitaentry.NewRepository()
	qittaEntryList, err := qiitaEntryRepo.FindByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		return nil, err
	}
	return qittaEntryList, nil
}
