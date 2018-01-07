package qiitaentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
	"github.com/fujimisakari/go-crawler/backend/infrastracture/db"
)

type DAO interface {
	SelectByID(id int) (map[string]interface{}, error)
	SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error)
	BulkInsert(qiitaEntryList *entity.QiitaEntryList) error
	DeleteByCrawlEntryID(crawlEntryID int) error
}

func newDAO() DAO {
	return &db.QiitaEntryDAO{Cnn: db.SharedInstance().Connection}
}
