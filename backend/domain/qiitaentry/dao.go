package qiitaentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

type QiitaEntryDAO interface {
	SelectByID(id int) (map[string]interface{}, error)
	SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error)
	BulkInsert(qiitaEntryList *entity.QiitaEntryList) error
	DeleteByCrawlEntryID(crawlEntryID int) error
}
