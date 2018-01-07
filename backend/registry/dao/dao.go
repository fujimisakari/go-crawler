package dao

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
	"github.com/fujimisakari/go-crawler/backend/infrastracture/db"
)

type DAO interface {
	CrawlEntry() crawlentry.CrawlEntryDAO
	HatenaHotEntry() hatenahotentry.HatenaHotEntryDAO
	QiitaEntry() qiitaentry.QiitaEntryDAO
}

func New() DAO {
	return daoImpl{}
}

type daoImpl struct{}

func (d daoImpl) CrawlEntry() crawlentry.CrawlEntryDAO {
	return &db.CrawlEntryDAO{Cnn: db.SharedInstance().Connection}
}

func (d daoImpl) HatenaHotEntry() hatenahotentry.HatenaHotEntryDAO {
	return &db.HatenaHotEntryDAO{Cnn: db.SharedInstance().Connection}
}

func (d daoImpl) QiitaEntry() qiitaentry.QiitaEntryDAO {
	return &db.QiitaEntryDAO{Cnn: db.SharedInstance().Connection}
}
