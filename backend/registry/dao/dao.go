package dao

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawl_entry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatena_hotentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiita_entry"
	"github.com/fujimisakari/go-crawler/backend/infrastracture/db"
)

type DAO interface {
	CrawlEntry() crawl_entry.CrawlEntryDAO
	HatenaHotEntry() hatena_hotentry.HatenaHotEntryDAO
	QiitaEntry() qiita_entry.QiitaEntryDAO
}

func New() DAO {
	return daoImpl{}
}

type daoImpl struct{}

func (d daoImpl) CrawlEntry() crawl_entry.CrawlEntryDAO {
	return &db.CrawlEntryDAO{Cnn: db.SharedInstance().Connection}
}

func (d daoImpl) HatenaHotEntry() hatena_hotentry.HatenaHotEntryDAO {
	return &db.HatenaHotEntryDAO{Cnn: db.SharedInstance().Connection}
}

func (d daoImpl) QiitaEntry() qiita_entry.QiitaEntryDAO {
	return &db.QiitaEntryDAO{Cnn: db.SharedInstance().Connection}
}
