package hatenahotentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
	"github.com/fujimisakari/go-crawler/backend/infrastracture/db"
)

type DAO interface {
	SelectByID(id int) (map[string]interface{}, error)
	SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error)
	BulkInsert(hatenaHotEntryList *entity.HatenaHotEntryList) error
	DeleteByCrawlEntryID(crawlEntryID int) error
}

func newDAO() DAO {
	return &db.HatenaHotEntryDAO{Cnn: db.SharedInstance().Connection}
}
